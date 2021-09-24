package doctors

import (
	"backend/business/doctors"
	"backend/controllers"
	"backend/controllers/doctors/requests"
	"backend/controllers/doctors/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DoctorController struct {
	DoctorUseCase doctors.Usecase
}

func NewDoctorController(doctorUseCase doctors.Usecase) *DoctorController {
	return &DoctorController{
		DoctorUseCase: doctorUseCase,
	}
}

func (doctorController DoctorController) Login(c echo.Context) error {
	doctorLogin := requests.Doctorlogin{}
	c.Bind(&doctorLogin)

	ctx := c.Request().Context()
	doctor, err := doctorController.DoctorUseCase.Login(ctx, doctorLogin.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(doctor))

}

func (doctorController DoctorController) Update(c echo.Context) error {
	DoctorUpdate := requests.DoctorUpdate{}
	c.Bind(&DoctorUpdate)

	ctx := c.Request().Context()
	doctor, err := doctorController.DoctorUseCase.Update(ctx, DoctorUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(doctor))
}

func (doctorController DoctorController) Register(c echo.Context) error {
	DoctorRegister := requests.DoctorRegister{}
	c.Bind(&DoctorRegister)

	ctx := c.Request().Context()
	doctor, err := doctorController.DoctorUseCase.Register(ctx, DoctorRegister.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(doctor))
}
