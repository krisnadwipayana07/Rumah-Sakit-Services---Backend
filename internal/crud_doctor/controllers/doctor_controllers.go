package controllers

import (
	"backend/configs"
	"backend/internal/crud_doctor/models"
	"backend/pkg/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InsertDoctor(c echo.Context) error {
	var dataDoctor models.DoctorRegister

	c.Bind(dataDoctor)

	//validate
	if dataDoctor.Name == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.NIP == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "NIP Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.Bidang == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bidang Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.ContactPerson == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Contact Person Still Empty",
			Data:    nil,
		})
	}

	var doctorDB database.Doctors
	doctorDB.Name = dataDoctor.Name
	doctorDB.NIP = dataDoctor.NIP
	doctorDB.Bidang = dataDoctor.Bidang
	doctorDB.ContactPerson = dataDoctor.ContactPerson

	result := configs.DB.Create(&doctorDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, configs.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, configs.BaseResponse{
		Code:    http.StatusOK,
		Message: "Inserted Doctors Data ",
		Data:    doctorDB,
	})

}
