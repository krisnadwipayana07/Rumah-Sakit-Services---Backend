package requests

import "backend/business/visitors"

type ShowRequest struct {
	SchedulesId uint `json:"schedulesId"`
	PatientsId  uint `json:"patientsId"`
}

func (showReq *ShowRequest) ToDomain() visitors.Domain {
	return visitors.Domain{
		SchedulesId: showReq.SchedulesId,
		PatientsId:  showReq.PatientsId,
	}
}
