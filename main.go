package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	ProfilePicUrl string `json:"profile_pic_url"`
}

var u1 = User{
	Username: "golang",
	Fullname: "Gopher",
	ProfilePicUrl: "https://golang.org/doc/gopher/frontpage.png",
}
var u2 = User{
	Username: "golang-way",
	Fullname: "go way",
	ProfilePicUrl: "https://golang.org/doc/gopher/frontpage.png",
}
var u3 = User{
	Username: "python",
	Fullname: "Pythonistas",
	ProfilePicUrl: "https://golang.org/doc/gopher/frontpage.png",
}
var u4 = User{
	Username: "Javascript",
	Fullname: "Use me for everything",
	ProfilePicUrl: "https://golang.org/doc/gopher/frontpage.png",
}

var allUsers = []User{
	u1, u2, u3, u4,
}

func usersHandler(c echo.Context) error {
	f := c.QueryParam("fileter")
	if f == "" {
		err := c.JSON(http.StatusOK, allUsers)
		return err
	}

	var users []User
	for _, u := range allUsers {
		if strings.Contains(u.Username, f) {
			users = append(users, u)
		}

		// if user.Username == f {
		// 	users = append(users, user)
		// }
	}

	return c.JSON(http.StatusOK, users)
}


func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/users", usersHandler)

	port := "2021"
	fmt.Println("starting ... port:", port)
	log.Println("starting ... port:", port)
	log.Fatal(e.Start(":" + port))
}
