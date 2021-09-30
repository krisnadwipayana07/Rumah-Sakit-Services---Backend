package requests

import "backend/business/visitors"

type AddRequest struct {
	SchedulesId uint   `json:"schedulesId"`
	PatientsId  uint   `json:"patientsId"`
	Keluhan     string `json:"keluhan"`
}

func (addReq *AddRequest) ToDomain() visitors.Domain {
	return visitors.Domain{
		SchedulesId: addReq.SchedulesId,
		PatientsId:  addReq.PatientsId,
		Keluhan:     addReq.Keluhan,
	}
}
