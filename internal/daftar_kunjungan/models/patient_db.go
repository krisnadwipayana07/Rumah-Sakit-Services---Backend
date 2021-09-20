package models

import "time"

type Patient struct {
	Id       int       `gorm:"primaryKey" json:"id"`
	Name     string    `json : "name"`
	Email    string    `json : "email" gorm:unique`
	Address  string    `json : "address"`
	NoBPJS   string    `json : "noBPJS" gorm:unique`
	CreateAt time.Time `json : "createAt"`
	UpdateAt time.Time `json : "updateAt"`
}
