package main

import (
	"backend/mvc/configs"
	"backend/mvc/routes"
	"log"
	"net/http"
)

func main() {
	configs.InitDB()
	connectLocalhost()
}

func connectLocalhost() {
	e := routes.AllRoutes()
	if err := e.Start(":8001"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
