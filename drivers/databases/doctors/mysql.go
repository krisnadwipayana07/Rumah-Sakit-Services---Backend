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

	raw := FromDomain(domain)
	raw.UpdateAt = time.Now()

	result := rep.Conn.Model(&doctor).Where("id = ?", raw.ID).Updates(raw)
	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return raw.ToDomain(), nil
}

func (rep *MysqlDoctorRepository) Register(ctx context.Context, domain doctors.Domain) (doctors.Domain, error) {
	doctorInsert := FromDomain(domain)
	doctorInsert.CreatedAt = time.Now()

	result := rep.Conn.Create(&doctorInsert)
	if result.Error != nil {
		return doctors.Domain{}, result.Error
	}

	return doctorInsert.ToDomain(), nil
}

// func (rep *MysqlDoctorRepository) AddSchedule(ctx context.Context, domain doctors.Domain, doctorId, scheduleId uint) ([]uint, error) {
// 	resultDoctor := FromDomain(domain)

// 	resultSchedule := rep.Conn.Raw("SELECT * FROM `doctors` WHERE id = ?", doctorId).Save(&domain.Schedules)
// 	if resultSchedule.Error != nil {
// 		return []uint{}, resultSchedule.Error
// 	}

// 	result := rep.Conn.Model(&resultDoctor).Association("Schedules").Append(domain.Schedules)
// 	if result != nil {
// 		return []uint{}, result
// 	}

// 	return []uint{doctorId, scheduleId}, nil
// }

// func (rep *MysqlDoctorRepository) AddSchedule(ctx context.Context, domain doctors.Domain, doctorId, scheduleId uint) ([]uint, error) {
// 	type Relation struct {
// 		DoctorsId   uint
// 		SchedulesId uint
// 	}
// 	Relation.DoctorsId = doctorId
// 	Relation.SchedulesId = scheduleId
// 	result := rep.Conn.Model(&domain).Association("doctor_schedules").Append()
// 	if result != nil {
// 		return []uint{}, result
// 	}

// 	return []uint{doctorId, scheduleId}, nil
// }
