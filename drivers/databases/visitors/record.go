package visitors

import (
	"backend/business/visitors"
	"backend/drivers/databases/patients"
	"backend/drivers/databases/schedules"
	"time"
)

type Visitors struct {
	// ID          uint `gorm:"primaryKey"`
	// DoctorId   uint `gorm:"primaryKey"` // --  docktor
	SchedulesId uint `gorm:"primaryKey"` //object shcedule
	PatientsId  uint `gorm:"primaryKey"` // patients
	AntrianId   uint
	Keluhan     string
	Patient     patients.Patients `gorm:"foreignKey:PatientsId"`
	CreateAt    time.Time
	UpdateAt    time.Time
}
type VisitorsLog struct {
	// ID          uint `gorm:"primaryKey"`
	// DoctorId   uint `gorm:"primaryKey"` // --  docktor
	SchedulesId uint
	PatientsId  uint
	AntrianId   uint
	Keluhan     string
	Solusi      string
	Message     string
	Schedule    schedules.Schedules `gorm:"foreignKey:SchedulesId"`
	Patients    patients.Patients   `gorm:"foreignKey:PatientsId"`
	CreateAt    time.Time
}

func ToDomainList(record []Visitors) []visitors.Domain {
	var returnValue []visitors.Domain
	for _, value := range record {
		returnValue = append(returnValue, value.ToDomain())
	}
	return returnValue
}

func (visitor *Visitors) ToDomain() visitors.Domain {
	return visitors.Domain{
		// ID:         visitor.ID,
		AntrianId:   visitor.AntrianId,
		SchedulesId: visitor.SchedulesId,
		PatientsId:  visitor.PatientsId,
		Keluhan:     visitor.Keluhan,
		Patient:     visitor.Patient.ToDomain(),
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

func ToDomainListLog(record []VisitorsLog) []visitors.Log {
	var returnValue []visitors.Log
	for _, value := range record {
		returnValue = append(returnValue, value.ToDomainLog())
	}
	return returnValue
}

func (visitor *VisitorsLog) ToDomainLog() visitors.Log {
	return visitors.Log{
		AntrianId:   visitor.AntrianId,
		SchedulesId: visitor.SchedulesId,
		PatientsId:  visitor.PatientsId,
		Keluhan:     visitor.Keluhan,
		Solusi:      visitor.Solusi,
		CreateAt:    visitor.CreateAt,
	}
}
func FromDomainLog(domain visitors.Log) VisitorsLog {
	return VisitorsLog{
		AntrianId:   domain.AntrianId,
		SchedulesId: domain.SchedulesId,
		PatientsId:  domain.PatientsId,
		Keluhan:     domain.Keluhan,
		Solusi:      domain.Solusi,
		Message:     domain.Message,
		CreateAt:    domain.CreateAt,
	}
}

func FromVisitorToLog(visitor Visitors) VisitorsLog {
	return VisitorsLog{
		AntrianId:   visitor.AntrianId,
		SchedulesId: visitor.SchedulesId,
		PatientsId:  visitor.PatientsId,
		Keluhan:     visitor.Keluhan,
		CreateAt:    visitor.CreateAt,
	}
}
