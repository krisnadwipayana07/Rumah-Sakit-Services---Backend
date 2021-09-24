package requests

import "backend/business/doctors"

type Doctorlogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (doctorLogin *Doctorlogin) ToDomain() doctors.Domain {
	return doctors.Domain{
		Email:    doctorLogin.Email,
		Password: doctorLogin.Password,
	}
}
