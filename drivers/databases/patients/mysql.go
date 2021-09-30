package patients

import (
	"backend/business/patients"
	"context"
	"time"

	"gorm.io/gorm"
)

type MysqlPatientRepository struct {
	Conn *gorm.DB
}

func NewMysqlPatientRepository(conn *gorm.DB) patients.Repository {
	return &MysqlPatientRepository{
		Conn: conn,
	}
}

func (rep *MysqlPatientRepository) Login(ctx context.Context, domain patients.Domain) (patients.Domain, error) {
	var patient Patients
	result := rep.Conn.First(&patient, "email = ? AND password = ?", domain.Email, domain.Password)

	if result.Error != nil {
		return patients.Domain{}, result.Error
	}

	return patient.ToDomain(), nil
}

func (rep *MysqlPatientRepository) Update(ctx context.Context, domain patients.Domain) (patients.Domain, error) {
	var patient Patients

	raw := FromDomain(domain)
	raw.UpdatedAt = time.Now()

	result := rep.Conn.Model(&patient).Where("id = ?", raw.ID).Updates(raw)
	if result.Error != nil {
		return patients.Domain{}, result.Error
	}

	return raw.ToDomain(), nil
}

func (rep *MysqlPatientRepository) Register(ctx context.Context, domain patients.Domain) (patients.Domain, error) {
	patientsInsert := FromDomain(domain)
	patientsInsert.CreatedAt = time.Now()

	result := rep.Conn.Create(&patientsInsert)
	if result.Error != nil {
		return patients.Domain{}, result.Error
	}

	return patientsInsert.ToDomain(), nil
}
