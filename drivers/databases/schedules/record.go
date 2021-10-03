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
	Patients    []patients.Patients `gorm:"many2many:visitors;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Doctor      Doctors        `gorm:"foreignKey:DoctorId"`
	// Doctor    []doctors.Doctors `gorm:"foreignKey:DoctorId"`
}
type Doctors struct {
	ID            uint   `gorm:"primaryKey"`
	Email         string `gorm:"unique"`
	Name          string
	Nip           string
	Description   string
	DoctorJob     string
	ContactPerson string
	CreatedAt     time.Time
}

func ToDomainList(record []Schedules) []schedules.Domain {
	var returnValue []schedules.Domain
	for _, value := range record {
		returnValue = append(returnValue, value.ToDomain())
	}
	return returnValue
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
		Doctors: schedules.Doctors{
			ID:            schedule.DoctorId,
			Email:         schedule.Doctor.Email,
			Name:          schedule.Doctor.Name,
			Nip:           schedule.Doctor.Nip,
			Description:   schedule.Doctor.Description,
			DoctorJob:     schedule.Doctor.DoctorJob,
			ContactPerson: schedule.Doctor.ContactPerson,
		},
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
