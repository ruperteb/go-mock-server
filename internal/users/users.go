package users

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"

	"echoserver/internal/db"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(c echo.Context) error {

	var users []User
	db.DB.Find(&users)

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	lastUserID := func() int {
		if len(users) > 0 {
			return users[len(users)-1].ID
		}

		return 0
	}

	u := new(User)
	u.ID = lastUserID() + 1

	if err := c.Bind(u); err != nil {
		return err
	}

	db.DB.Create(u)

	return c.JSON(http.StatusCreated, u)

}

func GetUsers(c echo.Context) error {

	var users []User
	db.DB.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {

	params := c.QueryParams()

	id, idExists := params["id"]
	name, nameExists := params["name"]

	u := new(User)
	if idExists {
		db.DB.First(&u, id[0])
		return c.JSON(http.StatusOK, u)
	}
	if nameExists {
		db.DB.Where("name = ?", name[0]).First(&u)
		return c.JSON(http.StatusOK, u)
	}

	return c.JSON(http.StatusOK, "No user found")
}

func DeleteUser(c echo.Context) error {
	params := c.QueryParams()

	id, idExists := params["id"]

	if idExists {
		db.DB.Delete(&User{}, id[0])
		return c.JSON(http.StatusOK, fmt.Sprint("User deleted:", id[0]))
	}
	return c.JSON(http.StatusOK, "No user found")
}

func UpdateUser(c echo.Context) error {

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	db.DB.Save(u)

	return c.JSON(http.StatusCreated, u)

}
