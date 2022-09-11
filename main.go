package main

import (
	assetHandler "backend-core/asset/delivery"
	assetRepository "backend-core/asset/repository"
	assetUseCase "backend-core/asset/usecase"
	authJwtHttpHandler "backend-core/auth/delivery"
	authJwtUseCase "backend-core/auth/usecase"
	"backend-core/common"
	"backend-core/docs"
	"backend-core/domain"
	orderHttpHandler "backend-core/order/delivery"
	orderRepository "backend-core/order/repository"
	orderUseCase "backend-core/order/usecase"
	userPlanRequestHttpHandler "backend-core/plan-request/delivery"
	userPlanRequestRepository "backend-core/plan-request/repository"
	userPlanRequestUseCase "backend-core/plan-request/usecase"
	planHandler "backend-core/plan/delivery"
	planGormRepository "backend-core/plan/repository"
	planGormUseCase "backend-core/plan/usecase"
	profileUseHandler "backend-core/profile/delivery"
	profileRepository "backend-core/profile/repository"
	profileUseCase "backend-core/profile/usecase"
	transactionHandler "backend-core/transaction/delivery"
	transactionRepository "backend-core/transaction/repository"
	transactionUseCase "backend-core/transaction/usecase"
	userPlanHttpHnadler "backend-core/user-plan/delivery"
	userPlanRepository "backend-core/user-plan/repository"
	userPlanUseCase "backend-core/user-plan/usecase"
	userHttpHandler "backend-core/user/delivery"
	userGormRepository "backend-core/user/repository"
	userUseCase "backend-core/user/usecase"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/syslog"
	"net/http"
	"os"
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
	log.Error().Err(err)
	c.Logger().Error(err)
}

// @title    B2B API Documentation
// @version  1.0

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description                 Description for what is this security definition being used
func main() {
	w, err := syslog.Dial("udp", "logs5.papertrailapp.com:19181", syslog.LOG_EMERG|syslog.LOG_KERN, "myapp")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	wr := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, w)
	log.Logger = zerolog.New(wr)

	//if err != nil {
	//	log.Panic().Msg("failed to dial syslog")
	//}

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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	err = db.AutoMigrate(&domain.User{}, domain.VerificationCode{}, domain.Profile{}, domain.UserPlan{})
	err = db.AutoMigrate(&domain.Asset{}, domain.Plan{}, domain.Transaction{}, &domain.UserRole{})
	err = db.AutoMigrate(&domain.UserPlanTransaction{}, domain.Order{}, domain.PlanRequest{}, &domain.OrderStatus{}, &domain.RequestStatus{}, &domain.Request{}, domain.TransactionType{})

	if err != nil {
		log.Error().Err(err)
	}

	if err != nil {
		log.Error().Err(err)
	}

	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Use(common.CORSMiddleWare(), middleware.RequestID())

	//e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	//	LogURI:    true,
	//	LogStatus: true,
	//	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
	//		log.Info().
	//			Str("URI", v.URI).
	//			Int("status", v.Status).
	//			Str("method", v.Method).
	//			Str("request-id", v.RequestID).
	//			Msg("plan-request")
	//		return nil
	//	},
	//}))

	docs.SwaggerInfo.Host = os.Getenv("BASE_URL")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	profRepository := profileRepository.NewProfileRepository(db)
	asstRepository := assetRepository.NewAssetRepository(db)
	usrRepository := userGormRepository.NewGormUserRepository(db)
	transRepository := transactionRepository.NewTransactionRepository(db)
	plnRepository := planGormRepository.NewPlanGormRepository(db)
	usrplnRepository := userPlanRepository.NewUserPlanRepository(db)
	ordrRespository := orderRepository.NewOrderRepository(db)
	usrplnRequestRepository := userPlanRequestRepository.NewPlanRequestRepository(db)

	profUseCase := profileUseCase.NewProfileUseCase(profRepository)
	asstUseCase := assetUseCase.NewAssetUseCase(asstRepository)
	usrUseCase := userUseCase.NewUserUseCase(usrRepository, asstUseCase, profUseCase, db)
	plnUseCase := planGormUseCase.NewPlanUseCase(plnRepository)
	transUseCase := transactionUseCase.NewTransactionUseCase(transRepository, usrUseCase, asstUseCase, db)
	athUseCase := authJwtUseCase.NewJwtAuthUseCase(usrUseCase, profUseCase)
	usrPlanUseCase := userPlanUseCase.NewUserPlanUseCase(usrplnRepository)
	ordrUseCase := orderUseCase.NewOrderUseCase(ordrRespository, usrPlanUseCase)
	usrplnRequestUsecae := userPlanRequestUseCase.NewPlanRequestUseCase(usrplnRequestRepository)

	profileUseHandler.NewProfileHttpHandler(e, profUseCase, usrUseCase)
	userHttpHandler.NewUserHttpHandler(e, usrUseCase, plnUseCase)
	transactionHandler.NewTransactionHttpHandler(e, transUseCase)
	assetHandler.NewAssetHttpHandler(e, asstUseCase, transUseCase)
	planHandler.NewPlanHttpHandler(e, plnUseCase)
	authJwtHttpHandler.NewAuthHttpHandler(e, athUseCase)
	orderHttpHandler.NewOrderHttpHandler(e, ordrUseCase)
	userPlanHttpHnadler.NewUserPlanDelivery(e, usrPlanUseCase)
	userPlanRequestHttpHandler.NewPlanRequestHttpHandler(e, usrplnRequestUsecae)

	e.Logger.Fatal(e.Start(":8080"))
}
