package controllers

import (
	"backend/configs"
	"backend/internal/daftar_kunjungan/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TambahDataPatient(c echo.Context) error {
	var kunjunganRegister models.PatientRegister

	c.Bind(kunjunganRegister)

	//validasi per data
	if kunjunganRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name Still Empty",
			Data:    nil,
		})
	}
	if kunjunganRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email Still Empty",
			Data:    nil,
		})
	}
	if kunjunganRegister.Address == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Address Still Empty",
			Data:    nil,
		})
	}
	if kunjunganRegister.NoBPJS == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "NoBPJS Still Empty",
			Data:    nil,
		})
	}

	var userDB models.Patient
}
