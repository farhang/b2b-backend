package main

import (
	authJwtHttpHandler "backend-core/auth/delivery"
	authJwtUseCase "backend-core/auth/usecase"
	"backend-core/common"
	"backend-core/docs"
	"backend-core/domain"
	userHttpHandler "backend-core/user/delivery"
	userGormRepository "backend-core/user/repository"
	userUseCase "backend-core/user/usecase"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/swaggo/echo-swagger"
)

type DBConfig struct {
	Engine   string `json:"engine"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	SSlMode  string `json:"ssl_mode"`
	Log      bool   `json:"log"`
	TimeZone string `json:"time_zone"`
}
type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
type CommonErrorDTO struct {
	Errors  []ApiError  `json:"errors,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	hr, _ := err.(*echo.HTTPError)
	if hr != nil {
		message := hr.Message

		if err, ok := hr.Message.(error); ok {
			message = err.Error()
		}

		_ = c.JSON(hr.Code, &CommonErrorDTO{
			Errors:  nil,
			Message: message,
		})
		return
	}
	vr, _ := err.(validator.ValidationErrors)
	if vr != nil {
		var out []ApiError
		if errors.As(err, &vr) {
			out = make([]ApiError, len(vr))
			for i, fe := range vr {
				out[i] = ApiError{fe.Field(), fe.Tag()}
			}
		}

		_ = c.JSON(http.StatusUnprocessableEntity, CommonErrorDTO{Errors: out, Message: http.StatusText(http.StatusUnprocessableEntity)})
		return
	}
	_ = c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	c.Logger().Error(err)
}

// @title    Turkey Exchange API Documentation
// @version  1.0

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description                 Description for what is this security definition being used
func main() {
	config := DBConfig{
		Engine:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSlMode:  os.Getenv("DB_SSLMODE"),
		Log:      false,
		TimeZone: os.Getenv("Time_Zone"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s  dbname=%s  port=%s TimeZone=%s", config.Host, config.Username, config.Password, config.DBName, config.Port, config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&domain.User{}, domain.EmailVerification{}, domain.Profile{}, domain.Deposit{})
	err = db.AutoMigrate(&domain.Asset{}, domain.Plan{}, domain.Report{}, domain.Card{}, domain.Withdraw{}, domain.Credit{})

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Use(common.CORSMiddleWare())

	docs.SwaggerInfo.Host = os.Getenv("BASE_URL")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	usrRepository := userGormRepository.NewGormUserRepository(db)
	usrUseCase := userUseCase.NewUserUseCase(usrRepository)
	userHttpHandler.NewUserHttpHandler(e, usrUseCase)

	athUseCase := authJwtUseCase.NewJwtAuthUseCase(usrUseCase)
	authJwtHttpHandler.NewAuthHttpHandler(e, athUseCase)
	e.Logger.Fatal(e.Start(":8080"))
}
