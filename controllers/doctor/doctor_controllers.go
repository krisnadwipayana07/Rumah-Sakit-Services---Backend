package doctor

import (
	"backend/configs"
	"backend/models/doctors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InsertDoctor(c echo.Context) error {
	var dataDoctor doctors.DoctorRegister
	c.Bind(&dataDoctor)

	if dataDoctor.Email == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.Password == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.Name == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.Nip == "" {
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

	var doctorDB doctors.Doctors
	doctorDB.Email = dataDoctor.Email
	doctorDB.Password = dataDoctor.Password
	doctorDB.Name = dataDoctor.Name
	doctorDB.Nip = dataDoctor.Nip
	doctorDB.Bidang = dataDoctor.Bidang
	doctorDB.ContactPerson = dataDoctor.ContactPerson
	doctorDB.CreateAt = time.Now()

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
		Message: "Inserted Doctors Data",
		Data:    doctorDB,
	})

}

func GetAllDoctors(c echo.Context) error {
	doctors := []doctors.Doctors{}

	result := configs.DB.Find(&doctors)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNoContent, configs.BaseResponse{
				Code:    http.StatusNoContent,
				Message: "Database Kosong",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, configs.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, configs.BaseResponse{
		Code:    http.StatusOK,
		Message: "Get All data doctor Sucseed",
		Data:    doctors,
	})
}

func GetSpesificDoctors(c echo.Context) error {
	doctors := []doctors.Doctors{}
	id := c.Param("id")

	result := configs.DB.Where("id = ?", id).Find(&doctors)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNoContent, configs.BaseResponse{
				Code:    http.StatusNoContent,
				Message: "data tidak ditemukan",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, configs.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, configs.BaseResponse{
		Code:    http.StatusOK,
		Message: "Get data doctor sucsess",
		Data:    doctors,
	})
}

func DeleteDoctorData(c echo.Context) error {
	id := c.Param("id")
	findDelete := []doctors.Doctors{}

	configs.DB.Where("id = ?", id).Find(&findDelete)
	result := configs.DB.Where("id = ?", id).Delete(&doctors.Doctors{})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, configs.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
		})
	}
	if result.RowsAffected != 0 {
		return c.JSON(http.StatusOK, configs.BaseResponse{
			Code:    http.StatusOK,
			Message: "Data berhasil di delete",
			Data:    findDelete,
		})
	}
	return c.JSON(http.StatusNoContent, configs.BaseResponse{
		Code:    http.StatusNoContent,
		Message: "Data tidak ditemukan",
		Data:    nil,
	})
}

func UpdateDoctor(c echo.Context) error {
	var dataDoctor doctors.DoctorUpdate
	c.Bind(&dataDoctor)

	if dataDoctor.Id == 0 {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Id Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.Email == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.Password == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.Name == "" {
		return c.JSON(http.StatusBadRequest, configs.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name Still Empty",
			Data:    nil,
		})
	}
	if dataDoctor.Nip == "" {
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

	dataDoctor.UpdateAt = time.Now()

	configs.DB.Model(&doctors.Doctors{}).Where("id = ?", dataDoctor.Id).Updates(dataDoctor)

	return c.JSON(http.StatusOK, configs.BaseResponse{
		Code:    http.StatusOK,
		Message: "Data berhasil di ubah",
		Data:    dataDoctor,
	})

}
