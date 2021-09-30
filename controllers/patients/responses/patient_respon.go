package responses

import (
	"backend/business/patients"
	"time"
)

type PatientResponse struct {
	Id            uint      `json:"id"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	BirthDate     time.Time `json:"birthDate"`
	BirthPlace    string    `json:"birthPlace"`
	NoBPJS        string    `json:"noBPJS"`
	ContactPerson string    `json:"contactPerson"`
	CreatedAt     time.Time `json:"createAt"`
	UpdatedAt     time.Time `json:"updateAt"`
}

func FromDomain(domain patients.Domain) PatientResponse {
	return PatientResponse{
		Id:            domain.ID,
		Email:         domain.Email,
		Name:          domain.Name,
		Address:       domain.Address,
		BirthDate:     domain.BirthDate,
		BirthPlace:    domain.BirthPlace,
		NoBPJS:        domain.NoBPJS,
		ContactPerson: domain.ContactPerson,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
