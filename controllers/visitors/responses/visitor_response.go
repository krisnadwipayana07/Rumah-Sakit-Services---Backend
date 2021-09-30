package responses

import (
	"backend/business/visitors"
	"time"
)

type VisitorResponse struct {
	Id          uint      `json:"id"`
	AntreanId   uint      `json:"antreanId"`
	SchedulesId uint      `json:"schedulesId"`
	PatientsId  uint      `json:"patientsId"`
	Keluhan     string    `json:"keluhan"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

func FromDomain(domain visitors.Domain) VisitorResponse {
	return VisitorResponse{
		Id:          domain.ID,
		AntreanId:   domain.AntrianId,
		SchedulesId: domain.SchedulesId,
		PatientsId:  domain.PatientsId,
		Keluhan:     domain.Keluhan,
		CreateAt:    domain.CreateAt,
		UpdateAt:    domain.UpdateAt,
	}
}
