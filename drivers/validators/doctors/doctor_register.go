package doctors

import "backend/business/doctors"

type DoctorRegister struct {
	Email         string `validate:"required"`
	Password      string `validate:"required"`
	Name          string `validate:"required"`
	Address       string `validate:"required"`
	Nip           string `validate:"required"`
	DoctorJob     string `validate:"required"`
	ContactPerson string `validate:"required"`
}

func FromDomain(domain doctors.Domain) DoctorRegister {
	return DoctorRegister{
		Email:         domain.Email,
		Password:      domain.Password,
		Name:          domain.Name,
		Address:       domain.Address,
		Nip:           domain.Nip,
		DoctorJob:     domain.DoctorJob,
		ContactPerson: domain.ContactPerson,
	}
}
