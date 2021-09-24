package doctors

import (
	"backend/business/doctors"
	"context"
	"time"

	"gorm.io/gorm"
)

type MysqlDoctorRepository struct {
	Conn *gorm.DB
}

func NewMysqlDoctorRepository(conn *gorm.DB) doctors.Repository {
	return &MysqlDoctorRepository{
		Conn: conn,
	}
}

func (rep *MysqlDoctorRepository) Login(ctx context.Context, domain doctors.Domain) (doctors.Domain, error) {
	var doctor Doctors
	result := rep.Conn.First(&doctor, "email = ? AND password = ?", domain.Email, domain.Password)

	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return doctor.ToDomain(), nil
}

func (rep *MysqlDoctorRepository) Update(ctx context.Context, domain doctors.Domain) (doctors.Domain, error) {
	var doctor Doctors
	var doctorUpdate Doctors

	doctorUpdate.Name = domain.Name
	doctorUpdate.Address = domain.Address
	doctorUpdate.Nip = domain.Nip
	doctorUpdate.DoctorJob = doctor.DoctorJob
	doctorUpdate.Email = doctor.Email
	doctorUpdate.Description = doctor.Description
	doctorUpdate.ContactPerson = doctor.ContactPerson
	doctorUpdate.UpdateAt = time.Now()

	result := rep.Conn.Model(&doctor).Where("id = ?", domain.Id).Updates(doctorUpdate)
	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return doctor.ToDomain(), nil
}

func (rep *MysqlDoctorRepository) Register(ctx context.Context, domain doctors.Domain) (doctors.Domain, error) {
	var doctorInsert Doctors

	doctorInsert.Email = domain.Email
	doctorInsert.Password = domain.Password
	doctorInsert.Name = domain.Name
	doctorInsert.Address = domain.Address
	doctorInsert.Nip = domain.Nip
	doctorInsert.Description = domain.Description
	doctorInsert.DoctorJob = domain.DoctorJob
	doctorInsert.ContactPerson = domain.ContactPerson
	doctorInsert.CreateAt = time.Now()

	result := rep.Conn.Create(&doctorInsert)
	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return doctorInsert.ToDomain(), nil
}
