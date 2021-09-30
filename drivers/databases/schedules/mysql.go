package schedules

import (
	"backend/business/schedules"
	"context"
	"time"

	"gorm.io/gorm"
)

type MysqlSchedulesRepository struct {
	Conn *gorm.DB
}

func NewMysqlSchedulesRepository(conn *gorm.DB) schedules.Repository {
	return &MysqlSchedulesRepository{
		Conn: conn,
	}
}

func (rep *MysqlSchedulesRepository) Add(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
	newSchedule := FromDomain(domain)
	newSchedule.CreatedAt = time.Now()
	// var doctorDomain doctors.Doctors
	// doctorDomain.ID = newSchedule.DoctorId
	// doctorDomain.Schedules = append(doctorDomain.Schedules, newSchedule)

	// if domain.DoctorId != 0 {
	// 	result := rep.Conn.Create(&newSchedule).Association("doctors").Append(&doctors.Doctors{Id: domain.DoctorId})
	// 	if result != nil {
	// 		return schedules.Domain{}, result
	// 	}
	// 	return newSchedule.ToDomain(), nil
	// }
	result := rep.Conn.Create(&newSchedule)
	if result.Error != nil {
		return schedules.Domain{}, result.Error
	}
	return newSchedule.ToDomain(), nil
}

func (rep *MysqlSchedulesRepository) Remove(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
	deleteSquedule := FromDomain(domain)

	result := rep.Conn.Delete(&deleteSquedule)
	if result.Error != nil {
		return schedules.Domain{}, result.Error
	}

	return deleteSquedule.ToDomain(), nil
}

func (rep *MysqlSchedulesRepository) Modificate(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
	var schedule Schedules

	modificateSchedules := FromDomain(domain)
	modificateSchedules.UpdatedAt = time.Now()

	result := rep.Conn.Model(&schedule).Where("id = ?", modificateSchedules.ID).Updates(modificateSchedules)
	if result.Error != nil {
		return schedules.Domain{}, result.Error
	}
	return modificateSchedules.ToDomain(), nil
}

func (rep *MysqlSchedulesRepository) Show(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
	findSchedules := FromDomain(domain)

	result := rep.Conn.Find(&findSchedules).Where("id = ?", findSchedules.ID)
	if result.Error != nil {
		return schedules.Domain{}, result.Error
	}

	return findSchedules.ToDomain(), nil
}

func (rep *MysqlSchedulesRepository) GetAll(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
	record := FromDomain(domain)
	result := rep.Conn.Find(&record)
	if result.Error != nil {
		return schedules.Domain{}, result.Error
	}
	// fmt.Println(record)
	return record.ToDomain(), nil
}

// func (rep *MysqlSchedulesRepository) InsertDoctor(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
// 	getIdBoth := FromDomain(domain)
// 	result := rep.Conn.Raw("INSERT INTO `doctor_schedules` (doctors_id, schedules_id) VALUES (?, ?)", getIdBoth.DoctorId, getIdBoth.ID).Scan(&getIdBoth)
// 	if result.Error != nil {
// 		return schedules.Domain{}, result.Error
// 	}
// 	return getIdBoth.ToDomain(), nil
// }
