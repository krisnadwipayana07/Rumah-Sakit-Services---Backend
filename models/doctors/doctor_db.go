package doctors

import "time"

type Doctors struct {
	Id            int    `gorm : "primaryKey" json:"id"`
	Email         string `json : "email" gorm:"uniqueIndex"`
	Password      string `json : "password"`
	Name          string `json : "name"`
	Nip           string `json : "nip"`
	Bidang        string `json : "bidang"`
	ContactPerson string `json : "contactPerson"`
	CreateAt      time.Time
	UpdateAt      time.Time
}
