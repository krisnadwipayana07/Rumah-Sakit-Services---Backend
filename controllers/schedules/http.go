package schedules

import (
	"backend/business/schedules"
	"backend/controllers"
	"backend/controllers/schedules/requests"
	"backend/controllers/schedules/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ScheduleController struct {
	SchedulesUseCase schedules.Usecase
}

func NewScheduleController(scheduleUsecase schedules.Usecase) *ScheduleController {
	return &ScheduleController{
		SchedulesUseCase: scheduleUsecase,
	}
}

//tambah schedule
func (schedulesController ScheduleController) Add(c echo.Context) error {
	scheduleInsert := requests.AddRequest{}
	c.Bind(&scheduleInsert)

	//hanya bisa menerima date dengan format "YYYY-MM-DD"
	date, err := time.Parse("2006-01-02", scheduleInsert.TanggalJaga)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	schedule, err := schedulesController.SchedulesUseCase.Add(ctx, scheduleInsert.ToDomain(date))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(schedule))
}

//modifikasi schedule
func (schedulesController ScheduleController) Modif(c echo.Context) error {
	scheduleUpdate := requests.UpdateRequest{}
	c.Bind(&scheduleUpdate)

	//hanya bisa menerima date dengan format "YYYY-MM-DD"
	date, err := time.Parse("2006-01-02", scheduleUpdate.TanggalJaga)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	schedule, err := schedulesController.SchedulesUseCase.Modificate(ctx, scheduleUpdate.ToDomain(date))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(schedule))
}

//menampilkan satu schedule
func (schedulesController ScheduleController) Show(c echo.Context) error {
	getId := requests.ShowRequest{}
	c.Bind(&getId)

	ctx := c.Request().Context()
	data, err := schedulesController.SchedulesUseCase.Show(ctx, getId.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, data)

}

//hapus satu schedule (soft delete)
func (schedulesController ScheduleController) Remove(c echo.Context) error {
	getId := requests.ShowRequest{}
	c.Bind(&getId)

	ctx := c.Request().Context()
	data, err := schedulesController.SchedulesUseCase.Remove(ctx, getId.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, data)

}

func (schedulesController ScheduleController) GetAllInOneDoctor(c echo.Context) error {
	getId := requests.GetByDoctor{}
	c.Bind(&getId)

	ctx := c.Request().Context()
	data, err := schedulesController.SchedulesUseCase.GetAllInOneDoctor(ctx, getId.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	var returnValue []responses.ScheduleResponse
	for _, value := range data {
		returnValue = append(returnValue, responses.FromDomain(value))
	}
	return controllers.NewSuccesResponse(c, returnValue)
}

func (schedulesController ScheduleController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := schedulesController.SchedulesUseCase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	var returnValue []responses.ScheduleResponse
	for _, value := range data {
		returnValue = append(returnValue, responses.FromDomain(value))
	}
	return controllers.NewSuccesResponse(c, returnValue)
}

// func (schedulesController ScheduleController) InsertDoctor(c echo.Context) error {
// 	getBind := requests.InsertDoctor{}
// 	c.Bind(&getBind)

// 	ctx := c.Request().Context()
// 	data, err := schedulesController.SchedulesUseCase.InsertDoctor(ctx, getBind.ToDomain())
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controllers.NewSuccesResponse(c, data)
// }
