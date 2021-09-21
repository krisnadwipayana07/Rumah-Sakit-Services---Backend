package controllers

import "github.com/labstack/echo/v4"

func DoctorRoutes() {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	ev1.POST("add_doctor/", InsertDoctor)
}
