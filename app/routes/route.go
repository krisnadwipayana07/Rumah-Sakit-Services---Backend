package routes

import (
	"backend/controllers/doctors"
	"backend/controllers/patients"
	"backend/controllers/schedules"
	"backend/controllers/visitors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	DoctorController   doctors.DoctorController
	PatientController  patients.PatientController
	ScheduleController schedules.ScheduleController
	VisitorController  visitors.VisitorController
}

func (cl *ControllerList) DoctorRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	ev1 := e.Group("/api/v1/doctor")
	ev1.POST("/login", cl.DoctorController.Login)
	ev1.POST("/update", cl.DoctorController.Update)
	ev1.POST("/register", cl.DoctorController.Register)
	// ev1.POST("/addSchedule", cl.DoctorController.AddSchedule)

}

func (cl *ControllerList) PatientRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.BodyDump(middlewares.Log))
	ev1 := e.Group("/api/v1/patient")
	ev1.POST("/login", cl.PatientController.Login)
	ev1.POST("/update", cl.PatientController.Update)
	ev1.POST("/register", cl.PatientController.Register)
}

func (cl *ControllerList) ScheduleRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.BodyDump(middlewares.Log))
	ev1 := e.Group("/api/v1/schedule")
	ev1.POST("/add", cl.ScheduleController.Add)
	ev1.POST("/update", cl.ScheduleController.Modif)
	ev1.POST("/show", cl.ScheduleController.Show)
	ev1.POST("/delete", cl.ScheduleController.Remove)
	ev1.GET("/getAll", cl.ScheduleController.GetAll)
	// ev1.POST("/insertDoctor", cl.ScheduleController.InsertDoctor)
}

func (cl *ControllerList) VisitorRoute(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.BodyDump(middlewares.Log))
	ev1 := e.Group("/api/v1/visitor")
	ev1.POST("/add", cl.VisitorController.AddVisitor)
}
