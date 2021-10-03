package responses

import (
	_domainSchedule "backend/business/schedules"
	"backend/drivers/databases/patients"
	"time"
)

type ScheduleResponse struct {
	Id          uint                    `json:"id"`
	DoctorId    uint                    `json:"doctorId"`
	Room        string                  `json:"room"`
	Message     string                  `json:"message"`
	TanggalJaga time.Time               `json:"tanggalJaga"`
	JamAwal     string                  `json:"jamAwal"`
	JamAkhir    string                  `json:"jamAkhir"`
	Patients    []patients.Patients     `json:"patients"`
	CreatedAt   time.Time               `json:"createdAt"`
	UpdatedAt   time.Time               `json:"updatedAt"`
	Doctors     _domainSchedule.Doctors `json:"doctors"`
}

func FromDomain(domain _domainSchedule.Domain) ScheduleResponse {
	return ScheduleResponse{
		Id:          domain.ID,
		DoctorId:    domain.DoctorId,
		Room:        domain.Room,
		Message:     domain.Message,
		TanggalJaga: domain.TanggalJaga,
		JamAwal:     domain.JamAwal,
		JamAkhir:    domain.JamAkhir,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		Doctors:     domain.Doctors,
	}
}
