package models

type PatientRegister struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	antrean string `json : "name"`
	Email   string `json : "email" gorm:unique`
	Address string `json : "address"`
	NoBPJS  string `json : "noBPJS" gorm:unique`
	Keluhan string `json : "keluhan"`
}
