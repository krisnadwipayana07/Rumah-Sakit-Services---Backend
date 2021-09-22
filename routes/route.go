package routes

import (
	"backend/controllers/doctor"

	"github.com/labstack/echo/v4"
)

func AllRoutes() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	ev1.POST("add_doctor", doctor.InsertDoctor)
	ev1.GET("get_all_doctor", doctor.GetAllDoctors)
	ev1.POST("delete_doctor/:id", doctor.DeleteDoctorData)
	return e
}
