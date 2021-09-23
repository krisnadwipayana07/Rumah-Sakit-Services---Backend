package routes

import (
	"backend/controllers/doctors"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	DoctorController doctors.DoctorController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	ev1 := e.Group("/api/v1")
	ev1.POST("/doctor/login", cl.DoctorController.Login)
	ev1.POST("/doctor/update", cl.DoctorController.Update)
	ev1.POST("/doctor/insert", cl.DoctorController.Register)
}
