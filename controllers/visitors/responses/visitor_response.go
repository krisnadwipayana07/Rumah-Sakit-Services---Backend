package responses

import (
	"backend/business/visitors"
	"time"
)

type VisitorResponse struct {
	// Id          uint      `json:"id"`
	AntreanId   uint      `json:"antreanId"`
	SchedulesId uint      `json:"schedulesId"`
	PatientsId  uint      `json:"patientsId"`
	Keluhan     string    `json:"keluhan"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}
type VisitorLogResponse struct {
	// Id          uint      `json:"id"`
	AntreanId   uint      `json:"antreanId"`
	SchedulesId uint      `json:"schedulesId"`
	PatientsId  uint      `json:"patientsId"`
	Keluhan     string    `json:"keluhan"`
	Solusi      string    `json:"solusi"`
	Message     string    `json:"message"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

func FromDomain(domain visitors.Domain) VisitorResponse {
	return VisitorResponse{
		// Id:          domain.ID,
		AntreanId:   domain.AntrianId,
		SchedulesId: domain.SchedulesId,
		PatientsId:  domain.PatientsId,
		Keluhan:     domain.Keluhan,
		CreateAt:    domain.CreateAt,
		UpdateAt:    domain.UpdateAt,
	}
}
func FromVisitorLog(domain visitors.Log) VisitorLogResponse {
	return VisitorLogResponse{
		// Id:          domain.ID,
		AntreanId:   domain.AntrianId,
		SchedulesId: domain.SchedulesId,
		PatientsId:  domain.PatientsId,
		Keluhan:     domain.Keluhan,
		Solusi:      domain.Solusi,
		Message:     domain.Message,
		CreateAt:    domain.CreateAt,
	}
}
