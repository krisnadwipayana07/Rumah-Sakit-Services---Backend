package responses

import (
	"backend/business/schedules"
	"time"
)

type ScheduleResponse struct {
	Id          uint      `json:"id"`
	DoctorId    uint      `json:"doctorId"`
	Room        string    `json:"room"`
	Message     string    `json:"message"`
	TanggalJaga time.Time `json:"tanggalJaga"`
	JamAwal     string    `json:"jamAwal"`
	JamAkhir    string    `json:"jamAkhir"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain schedules.Domain) ScheduleResponse {
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
	}
}
