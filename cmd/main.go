package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/user", show)
	e.POST("/user", show)
	e.GET("/mes/:id", mes)
	// e.POST("/users", display)

	e.Logger.Fatal(e.Start(":8080"))
}

func mes(c echo.Context) error {
	name := c.Param("id")
	return c.String(http.StatusOK, name)
}

func show(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	// return c.JSON(http.StatusOK, u)
	return c.JSON(http.StatusOK, test("ddd", "eee"))
}

func test(arg1, arg2 string) string {
	const xen = "xenTest"
	var arr = [...]string{"アイウエオ", "b"}

	for key, value := range arr {
		fmt.Println(key, value)
	}

	return arg1 + xen + arg2
}
