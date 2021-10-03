package doctors

import (
	"backend/business/doctors"
	"backend/drivers/databases/schedules"
	"time"

	"gorm.io/gorm"
)

type Doctors struct {
	ID            uint   `gorm:"primaryKey"`
	Email         string `gorm:"unique"`
	Password      string
	Name          string
	Address       string
	Nip           string
	Description   string
	DoctorJob     string
	Token         string
	ContactPerson string
	Schedule      []schedules.Schedules `gorm:"foreignKey:DoctorId"`
	CreatedAt     time.Time
	UpdateAt      time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	// Visitors		[]visitors.Visitors `gorm:"foreignKey:DoctorId;references:ID"`
	// Schedules []schedules.Schedules `gorm:"many2many:doctor_schedules"`
	// Visitors  []patients.Patients `gorm:"many2many:visitors;foreignKey:ID;joinForeignKey:DoctorId;References:ID;joinReferences:PatientId"`
}

func (doctor *Doctors) ToDomain() doctors.Domain {
	return doctors.Domain{
		ID:            doctor.ID,
		Email:         doctor.Email,
		Password:      doctor.Password,
		Name:          doctor.Name,
		Address:       doctor.Address,
		Nip:           doctor.Nip,
		Description:   doctor.Description,
		DoctorJob:     doctor.DoctorJob,
		Token:         doctor.Token,
		ContactPerson: doctor.ContactPerson,
		CreatedAt:     doctor.CreatedAt,
		UpdatedAt:     doctor.UpdateAt,
		// Schedules:     doctor.ToDomain().Schedules,
	}
}

func FromDomain(domain doctors.Domain) Doctors {
	return Doctors{
		ID:            domain.ID,
		Email:         domain.Email,
		Password:      domain.Password,
		Name:          domain.Name,
		Address:       domain.Address,
		Nip:           domain.Nip,
		Description:   domain.Description,
		DoctorJob:     domain.DoctorJob,
		Token:         domain.Token,
		ContactPerson: domain.ContactPerson,
		CreatedAt:     domain.CreatedAt,
		UpdateAt:      domain.UpdatedAt,
		// Schedules:     []schedules.Schedules{},
	}
}

func ToDomainList(record []Doctors) []doctors.Domain {
	var returnValue []doctors.Domain
	for _, value := range record {
		returnValue = append(returnValue, value.ToDomain())
	}
	return returnValue
}
