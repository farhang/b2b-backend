package common

import (
	"backend-core/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func AuthMiddleWare() echo.MiddlewareFunc {
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SuccessHandler: func(ctx echo.Context) {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(*domain.JwtCustomClaims)
			ctx.Set("userID", claims.UserId)
		},
		SigningKey: secret,
		Claims:     &domain.JwtCustomClaims{},
	})
}

func CORSMiddleWare() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	})
}
