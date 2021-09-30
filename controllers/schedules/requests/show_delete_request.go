package requests

import "backend/business/schedules"

type ShowRequest struct {
	Id uint `json:"id"`
}

func (showDeleteRequest *ShowRequest) ToDomain() schedules.Domain {
	return schedules.Domain{
		ID: showDeleteRequest.Id,
	}
}
