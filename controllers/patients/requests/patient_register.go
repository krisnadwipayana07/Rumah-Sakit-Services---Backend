package requests

import (
	"backend/business/patients"
	"time"
)

type PatientRegister struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	BirthDate     string `json:"birthDate"`
	BirthPlace    string `json:"birthPlace"`
	NoBPJS        string `json:"noBPJS"`
	ContactPerson string `json:"contactPerson"`
}

//hanya bisa menerima date dengan format "YYYY-MM-DD"
func (patientRegister *PatientRegister) ToDomain(date time.Time) patients.Domain {
	return patients.Domain{
		Email:         patientRegister.Email,
		Password:      patientRegister.Password,
		Name:          patientRegister.Name,
		Address:       patientRegister.Address,
		BirthDate:     date,
		BirthPlace:    patientRegister.BirthPlace,
		NoBPJS:        patientRegister.NoBPJS,
		ContactPerson: patientRegister.ContactPerson,
	}
}
