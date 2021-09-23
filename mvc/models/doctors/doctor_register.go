package doctors

type DoctorRegister struct {
	Email         string `json : "email" gorm:"uniqueIndex"`
	Password      string `json : "password"`
	Name          string `json : "name"`
	Nip           string `json : "nip"`
	Bidang        string `json : "bidang"`
	ContactPerson string `json : "contactPerson"`
}
