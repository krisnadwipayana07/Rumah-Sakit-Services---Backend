package visitors

import (
	"backend/business/visitors"
	"backend/controllers"
	"backend/controllers/visitors/requests"
	"backend/controllers/visitors/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VisitorController struct {
	VisitorUsecase visitors.Usecase
}

func NewVisitorController(visitControl visitors.Usecase) *VisitorController {
	return &VisitorController{
		VisitorUsecase: visitControl,
	}
}

func (visitorController VisitorController) AddVisitor(c echo.Context) error {
	ctx := c.Request().Context()

	req := requests.AddRequest{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := visitorController.VisitorUsecase.AddVisitor(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (visitorController VisitorController) UpdateVisitor(c echo.Context) error {
	ctx := c.Request().Context()

	req := requests.UpdateRequest{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := visitorController.VisitorUsecase.ModificateVisitor(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (visitorController VisitorController) DeleteVisitor(c echo.Context) error {
	ctx := c.Request().Context()

	req := requests.DeleteToLogRequest{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := visitorController.VisitorUsecase.RemoveVisitorToLog(ctx, req.ToDomainLog())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromVisitorLog(data))
}

func (visitorController VisitorController) ShowVisitor(c echo.Context) error {
	ctx := c.Request().Context()

	req := requests.ShowRequest{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := visitorController.VisitorUsecase.ShowVisitor(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}
