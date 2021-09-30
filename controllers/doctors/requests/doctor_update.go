package requests

import "backend/business/doctors"

type DoctorUpdate struct {
	Id            uint   `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Nip           string `json:"nip"`
	DoctorJob     string `json:"doctorJob"`
	Token         string `json:"token"`
	Description   string `json:"description"`
	ContactPerson string `json:"contactPerson"`
}

func (doctorUpdate *DoctorUpdate) ToDomain() doctors.Domain {
	return doctors.Domain{
		ID:            doctorUpdate.Id,
		Email:         doctorUpdate.Email,
		Password:      doctorUpdate.Password,
		Name:          doctorUpdate.Name,
		Address:       doctorUpdate.Address,
		Nip:           doctorUpdate.Nip,
		DoctorJob:     doctorUpdate.DoctorJob,
		Token:         doctorUpdate.Token,
		Description:   doctorUpdate.Description,
		ContactPerson: doctorUpdate.ContactPerson,
	}
}
