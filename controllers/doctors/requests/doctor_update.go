package requests

type DoctorUpdate struct {
	Id            int    `json:"id"`
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
