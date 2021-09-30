package routes

import (
	"backend/mvc/controllers/doctor"

	"github.com/labstack/echo/v4"
)

func AllRoutes() *echo.Echo {
	e := echo.New()
	eDoctor := e.Group("api/v1/doctor/")
	eDoctor.POST("add", doctor.InsertDoctor)
	eDoctor.GET("get_all", doctor.GetAllDoctors)
	eDoctor.GET("delete/:id", doctor.DeleteDoctorData)
	eDoctor.GET("spesific/:id", doctor.GetSpesificDoctors)
	eDoctor.POST("update_data", doctor.UpdateDoctor)
	return e
}
