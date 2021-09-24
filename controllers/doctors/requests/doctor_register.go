package requests

import "backend/business/doctors"

type DoctorRegister struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Nip           string `json:"nip"`
	Description   string `json:"description"`
	DoctorJob     string `json:"doctorJob"`
	ContactPerson string `json:"contactPerson"`
}

func (DoctorRegister *DoctorRegister) ToDomain() doctors.Domain {
	return doctors.Domain{
		Email:         DoctorRegister.Email,
		Password:      DoctorRegister.Password,
		Name:          DoctorRegister.Name,
		Address:       DoctorRegister.Name,
		Nip:           DoctorRegister.Nip,
		Description:   DoctorRegister.Description,
		DoctorJob:     DoctorRegister.DoctorJob,
		ContactPerson: DoctorRegister.ContactPerson,
	}
}
