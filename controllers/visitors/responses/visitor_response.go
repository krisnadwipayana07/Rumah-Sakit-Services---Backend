package responses

import (
	"backend/business/patients"
	"backend/business/visitors"
	"time"
)

type VisitorResponse struct {
	// Id          uint      `json:"id"`
	AntreanId   uint            `json:"antreanId"`
	SchedulesId uint            `json:"schedulesId"`
	PatientsId  uint            `json:"patientsId"`
	Keluhan     string          `json:"keluhan"`
	Patient     patients.Domain `json:"patient"`
	CreateAt    time.Time       `json:"createAt"`
	UpdateAt    time.Time       `json:"updateAt"`
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
		Patient: patients.Domain{
			ID:            domain.Patient.ID,
			Email:         domain.Patient.Email,
			Name:          domain.Patient.Name,
			Address:       domain.Patient.Address,
			BirthDate:     domain.Patient.BirthDate,
			BirthPlace:    domain.Patient.BirthPlace,
			NoBPJS:        domain.Patient.NoBPJS,
			ContactPerson: domain.Patient.ContactPerson,
		},
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
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
