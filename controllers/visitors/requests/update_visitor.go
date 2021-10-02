package requests

import "backend/business/visitors"

type UpdateRequest struct {
	AntrianId   uint   `json:"antrian_id"`
	SchedulesId uint   `json:"schedulesId"`
	PatientsId  uint   `json:"patientsId"`
	Keluhan     string `json:"keluhan"`
}

func (updateReq *UpdateRequest) ToDomain() visitors.Domain {
	return visitors.Domain{
		AntrianId:   updateReq.AntrianId,
		SchedulesId: updateReq.SchedulesId,
		PatientsId:  updateReq.PatientsId,
		Keluhan:     updateReq.Keluhan,
	}
}
