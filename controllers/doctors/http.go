package doctors

import (
	"backend/business/doctors"
	"backend/controllers"
	"backend/controllers/doctors/requests"
	"backend/controllers/doctors/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Doctor godoc
// @Summary doctor documentation
// @Description documentation of doctor
// @Tags doctor
// @Accept json
// @Produce json
// @Param todo body types.Todo true "New Todo"
// @Success 201 {object} responses.DoctorResponse
// @Failure 400 {object} controllers.NewErrorResponse
// @Router /todos [post]

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

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = doctor.Token
	c.SetCookie(cookie)

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

func (doctorController DoctorController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := doctorController.DoctorUseCase.ShowAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	var returnValue []responses.DoctorResponse
	for _, value := range data {
		returnValue = append(returnValue, responses.FromDomain(value))
	}
	return controllers.NewSuccesResponse(c, returnValue)
}

// func (doctorController DoctorController) AddSchedule(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	addSchedule := requests.AddScheduleDoctor{}
// 	if err := c.Bind(&addSchedule); err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	schedule, err := doctorController.DoctorUseCase.AddSchedule(ctx, addSchedule.DoctorId, addSchedule.ScheduleId)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccesResponse(c, requests.FromRequest(schedule[0], schedule[1]))
// }
