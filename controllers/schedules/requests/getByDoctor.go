package requests

import "backend/business/schedules"

type GetByDoctor struct {
	DoctorId uint `json:"doctorId"`
}

func (scheduleAdd *GetByDoctor) ToDomain() schedules.Domain {
	return schedules.Domain{
		DoctorId: scheduleAdd.DoctorId,
	}
}
