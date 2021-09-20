package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func getRunningAddress() string {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	return fmt.Sprintf("%s:%s", host, port)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "wellcome backend services")
	})
	// if e.Logger.Fatal(e.Start(":8001"))
	if err := e.Start(":8001"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	getRunningAddress()
}
