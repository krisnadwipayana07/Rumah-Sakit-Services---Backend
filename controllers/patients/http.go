package patients

import (
	"backend/business/patients"
	"backend/controllers"
	"backend/controllers/patients/requests"
	"backend/controllers/patients/responses"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

type PatientController struct {
	PatientUseCase patients.Usecase
}

func NewPatientController(patientUseCase patients.Usecase) *PatientController {
	return &PatientController{
		PatientUseCase: patientUseCase,
	}
}

func (patientController PatientController) Login(c echo.Context) error {
	patientLogin := requests.PatientsLogin{}
	c.Bind(&patientLogin)

	ctx := c.Request().Context()
	patient, err := patientController.PatientUseCase.Login(ctx, patientLogin.ToDomain())

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = patient.Token
	c.SetCookie(cookie)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(patient))

}

func (patientController PatientController) Update(c echo.Context) error {
	patientUpdate := requests.PatientUpdate{}
	c.Bind(&patientUpdate)

	//hanya bisa menerima date dengan format "YYYY-MM-DD"
	date, err := time.Parse("2006-01-02", patientUpdate.BirthDate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	patient, err := patientController.PatientUseCase.Update(ctx, patientUpdate.ToDomain(date))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(patient))
}

func (patientController PatientController) Register(c echo.Context) error {
	patientRegister := requests.PatientRegister{}
	c.Bind(&patientRegister)

	//hanya bisa menerima date dengan format "YYYY-MM-DD"
	date, err := time.Parse("2006-01-02", patientRegister.BirthDate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	patient, err := patientController.PatientUseCase.Register(ctx, patientRegister.ToDomain(date))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(patient))
}
