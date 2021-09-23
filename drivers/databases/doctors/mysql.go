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

func (rep *MysqlDoctorRepository) Login(ctx context.Context, email string, password string) (doctors.Domain, error) {
	var doctor Doctors
	result := rep.Conn.First(&doctor, "email = ? AND password = ?", email, password)

	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return doctor.ToDomain(), nil
}

func (rep *MysqlDoctorRepository) Update(ctx context.Context, id int, name, address, nip, doctorJob, email, description, contactPerson string) (doctors.Domain, error) {
	var doctor Doctors
	var doctorUpdate Doctors

	doctorUpdate.Name = name
	doctorUpdate.Address = address
	doctorUpdate.Nip = nip
	doctorUpdate.DoctorJob = doctorJob
	doctorUpdate.Email = email
	doctorUpdate.Description = description
	doctorUpdate.ContactPerson = contactPerson
	doctorUpdate.UpdateAt = time.Now()

	result := rep.Conn.Model(&doctor).Where("id = ?", id).Updates(doctorUpdate)
	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return doctor.ToDomain(), nil
}

func (rep *MysqlDoctorRepository) Register(ctx context.Context, email, password, name, nip, address, description, doctorJob, contactPerson string) (doctors.Domain, error) {
	var doctorInsert Doctors

	doctorInsert.Email = email
	doctorInsert.Password = password
	doctorInsert.Name = name
	doctorInsert.Address = address
	doctorInsert.Nip = nip
	doctorInsert.Description = description
	doctorInsert.DoctorJob = doctorJob
	doctorInsert.ContactPerson = contactPerson
	doctorInsert.CreateAt = time.Now()

	result := rep.Conn.Create(&doctorInsert)
	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return doctorInsert.ToDomain(), nil
}
