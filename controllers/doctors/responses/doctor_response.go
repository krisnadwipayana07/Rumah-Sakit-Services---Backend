package responses

import (
	"backend/business/doctors"
	"time"
)

type DoctorResponse struct {
	Id            int       `json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Nip           string    `json:"nip"`
	Description   string    `json:"description"`
	DoctorJob     string    `json:"doctorJob"`
	Token         string    `json:"token"`
	ContactPerson string    `json:"contactPerson"`
	CreateAt      time.Time `json:"createAt"`
	UpdateAt      time.Time `json:"updateAt"`
}

func FromDomain(domain doctors.Domain) DoctorResponse {
	return DoctorResponse{
		Id:            domain.Id,
		Email:         domain.Email,
		Password:      domain.Password,
		Name:          domain.Name,
		Address:       domain.Address,
		Nip:           domain.Nip,
		Description:   domain.Description,
		DoctorJob:     domain.DoctorJob,
		ContactPerson: domain.ContactPerson,
		CreateAt:      domain.CreateAt,
		UpdateAt:      domain.UpdateAt,
	}
}
