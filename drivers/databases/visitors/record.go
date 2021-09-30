package visitors

import (
	"backend/business/visitors"
	"time"
)

type Visitors struct {
	// ID          uint `gorm:"primaryKey"`
	SchedulesId uint `gorm:"primaryKey"` //object shcedule
	PatientsId  uint `gorm:"primaryKey"` // patients
	AntrianId   uint
	// DoctorId   uint `gorm:"primaryKey"` // --  docktor
	Keluhan  string
	CreateAt time.Time
	UpdateAt time.Time
}

func (visitor *Visitors) ToDomain() visitors.Domain {
	return visitors.Domain{
		// ID:         visitor.ID,
		AntrianId:   visitor.AntrianId,
		SchedulesId: visitor.SchedulesId,
		PatientsId:  visitor.PatientsId,
		Keluhan:     visitor.Keluhan,
		CreateAt:    visitor.CreateAt,
		UpdateAt:    visitor.UpdateAt,
	}
}
func FromDomain(domain visitors.Domain) Visitors {
	return Visitors{
		// ID:         domain.ID,
		AntrianId:   domain.AntrianId,
		SchedulesId: domain.SchedulesId,
		PatientsId:  domain.PatientsId,
		Keluhan:     domain.Keluhan,
		CreateAt:    domain.CreateAt,
		UpdateAt:    domain.UpdateAt,
	}
}
