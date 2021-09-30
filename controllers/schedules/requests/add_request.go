package requests

import (
	"backend/business/schedules"
	"time"
)

type AddRequest struct {
	DoctorId    uint   `json:"doctorId"`
	Room        string `json:"room"`
	Message     string `json:"message"`
	TanggalJaga string `json:"tanggalJaga"`
	JamAwal     string `json:"jamAwal"`
	JamAkhir    string `json:"jamAkhir"`
}

func (scheduleAdd *AddRequest) ToDomain(date time.Time) schedules.Domain {
	return schedules.Domain{
		DoctorId:    scheduleAdd.DoctorId,
		Room:        scheduleAdd.Room,
		Message:     scheduleAdd.Message,
		TanggalJaga: date,
		JamAwal:     scheduleAdd.JamAwal,
		JamAkhir:    scheduleAdd.JamAkhir,
	}
}
