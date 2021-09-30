package requests

import (
	"backend/business/schedules"
	"time"
)

type UpdateRequest struct {
	Id          uint   `json:"id"`
	DoctorId    uint   `json:"doctorId"`
	Room        string `json:"room"`
	Message     string `json:"message"`
	TanggalJaga string `json:"tanggalJaga"`
	JamAwal     string `json:"jamAwal"`
	JamAkhir    string `json:"jamAkhir"`
}

func (scheduleUpdate *UpdateRequest) ToDomain(date time.Time) schedules.Domain {
	return schedules.Domain{
		ID:          scheduleUpdate.Id,
		DoctorId:    scheduleUpdate.DoctorId,
		Room:        scheduleUpdate.Room,
		Message:     scheduleUpdate.Message,
		TanggalJaga: date,
		JamAwal:     scheduleUpdate.JamAwal,
		JamAkhir:    scheduleUpdate.JamAkhir,
	}
}
