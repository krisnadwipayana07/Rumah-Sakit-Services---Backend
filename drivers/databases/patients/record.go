package patients

import (
	"backend/business/patients"
	"time"

	"gorm.io/gorm"
)

type Patients struct {
	ID            uint   `gorm:"primaryKey"`
	Email         string `gorm:"unique"`
	Password      string
	Name          string
	Address       string
	BirthDate     time.Time `gorm:"DEFALUT"`
	BirthPlace    string
	NoBPJS        string
	Token         string
	ContactPerson string
	// Visitors      []visitors.Visitors `gorm:"foreignKey:PatientId;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (patient *Patients) ToDomain() patients.Domain {
	return patients.Domain{
		ID:            patient.ID,
		Email:         patient.Email,
		Password:      patient.Password,
		Name:          patient.Name,
		Address:       patient.Address,
		BirthDate:     patient.BirthDate,
		BirthPlace:    patient.BirthPlace,
		NoBPJS:        patient.NoBPJS,
		Token:         patient.Token,
		ContactPerson: patient.ContactPerson,
		CreatedAt:     patient.CreatedAt,
		UpdatedAt:     patient.UpdatedAt,
	}
}

func FromDomain(domain patients.Domain) Patients {
	return Patients{
		ID:            domain.ID,
		Email:         domain.Email,
		Password:      domain.Password,
		Name:          domain.Name,
		Address:       domain.Address,
		BirthDate:     domain.BirthDate,
		BirthPlace:    domain.BirthPlace,
		NoBPJS:        domain.NoBPJS,
		Token:         domain.Token,
		ContactPerson: domain.ContactPerson,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
