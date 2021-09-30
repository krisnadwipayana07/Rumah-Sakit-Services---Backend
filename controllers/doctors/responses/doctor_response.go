package responses

import (
	"backend/business/doctors"
	"backend/business/schedules"
	"time"
)

type DoctorResponse struct {
	Id            uint               `json:"id"`
	Email         string             `json:"email"`
	Password      string             `json:"password"`
	Name          string             `json:"name"`
	Address       string             `json:"address"`
	Nip           string             `json:"nip"`
	Description   string             `json:"description"`
	DoctorJob     string             `json:"doctorJob"`
	Token         string             `json:"token"`
	ContactPerson string             `json:"contactPerson"`
	Schedules     []schedules.Domain `json:"schedules"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
}

func FromDomain(domain doctors.Domain) DoctorResponse {
	return DoctorResponse{
		Id:            domain.ID,
		Email:         domain.Email,
		Password:      domain.Password,
		Name:          domain.Name,
		Address:       domain.Address,
		Nip:           domain.Nip,
		Description:   domain.Description,
		DoctorJob:     domain.DoctorJob,
		ContactPerson: domain.ContactPerson,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
