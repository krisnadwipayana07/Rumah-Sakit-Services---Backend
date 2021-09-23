package doctors

import (
	"backend/business/doctors"
	"time"
)

type Doctors struct {
	Id            int    `gorm:"primaryKey"`
	Email         string `gorm:"unique"`
	Password      string
	Name          string
	Address       string
	Nip           string
	Description   string
	DoctorJob     string
	Token         string
	ContactPerson string
	CreateAt      time.Time
	UpdateAt      time.Time
}

func (doctor *Doctors) ToDomain() doctors.Domain {
	return doctors.Domain{
		Id:            doctor.Id,
		Email:         doctor.Email,
		Password:      doctor.Password,
		Name:          doctor.Name,
		Address:       doctor.Address,
		Nip:           doctor.Nip,
		Description:   doctor.Description,
		DoctorJob:     doctor.DoctorJob,
		Token:         doctor.Token,
		ContactPerson: doctor.ContactPerson,
		CreateAt:      doctor.CreateAt,
		UpdateAt:      doctor.UpdateAt,
	}
}

func FromDomain(domain doctors.Domain) Doctors {
	return Doctors{
		Id:            domain.Id,
		Email:         domain.Email,
		Password:      domain.Password,
		Name:          domain.Name,
		Address:       domain.Address,
		Nip:           domain.Nip,
		Description:   domain.Description,
		DoctorJob:     domain.DoctorJob,
		Token:         domain.Token,
		ContactPerson: domain.ContactPerson,
		CreateAt:      domain.CreateAt,
		UpdateAt:      domain.UpdateAt,
	}
}
