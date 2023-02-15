package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	UserName string `json:"name"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type UserAuth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func GetToken(c echo.Context) error {
	u := new(UserAuth)

	if err := c.Bind(u); err != nil {
		return err
	}

	// Throws unauthorized error
	if u.UserName != "ruperteb@gmail.com" || u.Password != "test" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"ruperteb@gmail.com",
		"admin",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
