package requests

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
