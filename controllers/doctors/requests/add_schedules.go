package requests

type AddScheduleDoctor struct {
	DoctorId   uint `json:"doctorId"`
	ScheduleId uint `json:"scheduleId"`
}

func FromRequest(doctorId, scheduleId uint) AddScheduleDoctor {
	return AddScheduleDoctor{
		DoctorId:   doctorId,
		ScheduleId: scheduleId,
	}
}
