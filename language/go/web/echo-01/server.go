package main

import (
	"net/http"
	"io"
	"os"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func save(c echo.Context) error  {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func profile(c echo.Context) error  {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}

	defer dst.Close()
	if _, err = io.Copy(dst,src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you!" + name + "</b>")
}

func createUser(c echo.Context) error  {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Helllo")
	})

	e.GET("/users/:id", getUser)
	e.POST("/save", save)
	e.POST("/profile",profile)
	e.POST("/users", createUser)
	e.Logger.Fatal(e.Start(":9090"))

}
