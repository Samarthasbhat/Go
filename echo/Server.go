package main

import (
	"net/http"
	"os"
	"io"
	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Get user with ID: "+id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "Team: "+team+", Member: "+member)
}

func saveUser(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
  	avatar, err := c.FormFile("avatar")
  	if err != nil {
 		return err
 	}
 
 	// Source
 	src, err := avatar.Open()
 	if err != nil {
 		return err
 	}
 	defer src.Close()
 
 	// Destination
 	dst, err := os.Create(avatar.Filename)
 	if err != nil {
 		return err
 	}
 	defer dst.Close()
 
 	// Copy
 	if _, err = io.Copy(dst, src); err != nil {
  		return err
  	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/save", saveUser)
e.GET("/users/:id", getUser)
// e.PUT("/users/:id", updateUser)
// e.DELETE("/users/:id", deleteUser)
e.GET("/show", show)
	e.Logger.Fatal(e.Start(":8080"))
}