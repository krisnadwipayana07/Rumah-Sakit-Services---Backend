package patients

import "time"

type Patients struct {
	Id            int       `gorm : "primaryKey" json:"id"`
	Email         string    `json : "email" gorm:"uniqueIndex"`
	Password      string    `json : "password"`
	Name          string    `json : "name"`
	Alamat        string    `json : "alamat"`
	Bidang        string    `json : "bidang"`
	ContactPerson string    `json : "contactPerson"`
	CreateAt      time.Time `json : "createAt"`
	UpdateAt      time.Time `json : "updateAt"`
	DeleteAt      time.Time `json : "deleteAt"`
}
