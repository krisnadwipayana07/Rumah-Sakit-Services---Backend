package schedules

import (
	"backend/business/schedules"
	"backend/drivers/databases/patients"
	"time"

	"gorm.io/gorm"
)

type Schedules struct {
	ID          uint `gorm:"primaryKey"`
	DoctorId    uint
	Room        string
	Message     string
	TanggalJaga time.Time
	JamAwal     string
	JamAkhir    string
	// Doctor      []doctors.Doctors   `gorm:"foreignKey:DoctorId"`
	Visitor   []patients.Patients `gorm:"many2many:visitors;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (schedule *Schedules) ToDomain() schedules.Domain {
	return schedules.Domain{
		ID:          schedule.ID,
		DoctorId:    schedule.DoctorId,
		Room:        schedule.Room,
		Message:     schedule.Message,
		TanggalJaga: schedule.TanggalJaga,
		JamAwal:     schedule.JamAwal,
		JamAkhir:    schedule.JamAkhir,
		CreatedAt:   schedule.CreatedAt,
		UpdatedAt:   schedule.UpdatedAt,
	}
}

func FromDomain(domain schedules.Domain) Schedules {
	return Schedules{
		ID:          domain.ID,
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
