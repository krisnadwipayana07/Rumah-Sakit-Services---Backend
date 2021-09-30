package requests

import "backend/business/schedules"

type DeleteRequest struct {
	Id uint `json:"id"`
}

func (deleteRequest *DeleteRequest) ToDomain() schedules.Domain {
	return schedules.Domain{
		ID: deleteRequest.Id,
	}
}
