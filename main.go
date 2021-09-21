package main

import (
	"backend/configs"
	"backend/pkg/routes"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	routes.AllRoutes()
	connectLocalhost()
	configs.InitDB()
}

func connectLocalhost() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "wellcome backend services")
	})
	// if e.Logger.Fatal(e.Start(":8001"))
	if err := e.Start(":8001"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
