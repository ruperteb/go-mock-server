package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"fmt"

	"github.com/golang-jwt/jwt"

	"echoserver/internal/auth"
	"echoserver/internal/users"

	"echoserver/internal/db"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func CreateServer() {
	fmt.Println("test")

	e := echo.New()

	db.InitDB()

	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/getToken", auth.GetToken)

	authGroup := e.Group("/auth")
	authGroup.Use(middleware.JWTWithConfig(config))

	authGroup.GET("/getUser", users.GetUser)

	authGroup.POST("/createUser", users.CreateUser)

	authGroup.GET("/getUsers", users.GetUsers)

	authGroup.GET("/deleteUser", users.DeleteUser)

	authGroup.PUT("/updateUser", users.UpdateUser)

	e.Logger.Fatal(e.Start(":1323"))

}
