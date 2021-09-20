package models

type DoctorRegister struct {
	Name          string `json : "name`
	NIP           string
	Bidang        string `json : "bidang"`
	ContactPerson string `json : "contact_person"`
}
