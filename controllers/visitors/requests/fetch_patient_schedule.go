package requests

import "backend/business/visitors"

type GetAllPatient struct {
	SchedulesId uint `json:"schedulesId"`
}

func (getAll *GetAllPatient) ToDomain() visitors.Domain {
	return visitors.Domain{
		SchedulesId: getAll.SchedulesId,
	}
}
