package common

import (
	"backend-core/domain"
	"github.com/casbin/casbin/v2"
	"github.com/golang-jwt/jwt"
	casbin_mw "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func AuthMiddleWare() echo.MiddlewareFunc {
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SuccessHandler: func(ctx echo.Context) {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(*domain.JwtCustomClaims)
			log.Info().Int("user", claims.UserId).Msg("authenticated user")
			ctx.Set("userID", claims.UserId)
			ctx.Set("role", claims.Role)
		},
		SigningKey: secret,
		Claims:     &domain.JwtCustomClaims{},
	})
}
func CASBINMiddleWare() echo.MiddlewareFunc {
	ce, err := casbin.NewEnforcer("auth_model.conf", "policy.csv")
	if err != nil {
		log.Error().Err(err)
	}

	return casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
		Enforcer: ce,
		UserGetter: func(c echo.Context) (string, error) {
			//role := c.Get("role").(uint)
			return "MEMBER", nil
		},
	})
}

func CORSMiddleWare() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	})
}
