package models

import "time"

type Doctors struct {
	Id            int    `gorm : "primaryKey" json:"id"`
	Name          string `json : "name`
	NIP           string
	Bidang        string    `json : "bidang"`
	ContactPerson string    `json : "contact_person"`
	CreateAt      time.Time `json : "createAt"`
	UpdateAt      time.Time `json : "updateAt"`
}
