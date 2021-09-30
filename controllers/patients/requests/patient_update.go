package requests

import (
	"backend/business/patients"
	"time"
)

type PatientUpdate struct {
	Id            uint   `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	BirthDate     string `json:"birthDate"`
	BirthPlace    string `json:"birthPlace"`
	NoBPJS        string `json:"noBPJS"`
	Token         string `json:"token"`
	ContactPerson string `json:"contactPerson"`
}

//hanya bisa menerima date dengan format "YYYY-MM-DD"
func (patientUpdate *PatientUpdate) ToDomain(date time.Time) patients.Domain {
	return patients.Domain{
		ID:            patientUpdate.Id,
		Email:         patientUpdate.Email,
		Name:          patientUpdate.Name,
		Address:       patientUpdate.Address,
		BirthDate:     date,
		BirthPlace:    patientUpdate.BirthPlace,
		NoBPJS:        patientUpdate.NoBPJS,
		Token:         patientUpdate.Token,
		ContactPerson: patientUpdate.ContactPerson,
	}
}
