package requests

import (
	"backend/business/patients"
)

type PatientsLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (patientslogin *PatientsLogin) ToDomain() patients.Domain {
	return patients.Domain{
		Email:    patientslogin.Email,
		Password: patientslogin.Password,
	}
}
