package middlewares

import (
	"backend/controllers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCostumClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiredDuration int
}

func (config *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCostumClaims{},
		SigningKey: []byte(config.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controllers.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

func (config *ConfigJWT) GenerateTokenDoctor(userId uint) string {
	claims := JWTCostumClaims{
		userId,
		"doctor",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(config.ExpiredDuration))).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(config.SecretJWT))
	// if err != nil {
	// 	panic(err)
	// }
	return t
}

func (config *ConfigJWT) GenerateTokenPatient(userId uint) string {
	claims := JWTCostumClaims{
		userId,
		"patient",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(config.ExpiredDuration))).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(config.SecretJWT))
	// if err != nil {
	// 	panic(err)
	// }
	return t
}

func GetUser(c echo.Context) *JWTCostumClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCostumClaims)
	return claims
}
