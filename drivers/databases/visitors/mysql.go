package visitors

import (
	"backend/business/visitors"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MysqlVisitorRepository struct {
	Conn *gorm.DB
}

func NewMysqlVisitorRepository(conn *gorm.DB) visitors.Repository {
	return &MysqlVisitorRepository{
		Conn: conn,
	}
}

func (rep *MysqlVisitorRepository) AddVisitor(ctx context.Context, domain visitors.Domain) (visitors.Domain, error) {
	newVisitor := FromDomain(domain)
	fmt.Println(newVisitor.AntrianId)
	newVisitor.AntrianId = 1
	var result uint
	fmt.Println(newVisitor.AntrianId)

	hasilRaw := rep.Conn.Select("antrian_id").Table("visitors").Where("schedules_id = ?", newVisitor.SchedulesId).Order("schedules_id DESC").Find(&result)
	if hasilRaw.Error != nil {
		return visitors.Domain{}, hasilRaw.Error
	}
	fmt.Println(result)
	if result == 0 {
		newVisitor.AntrianId = 1
		fmt.Println("error catch")
	} else {
		newVisitor.AntrianId = result + 1
	}
	fmt.Println(newVisitor.AntrianId)
	newVisitor.CreateAt = time.Now()
	resultVisitor := rep.Conn.Create(&newVisitor)
	if resultVisitor.Error != nil {
		return visitors.Domain{}, resultVisitor.Error
	}
	return newVisitor.ToDomain(), nil
}

func (rep *MysqlVisitorRepository) RemoveVisitorToLog(ctx context.Context, log visitors.Log) (visitors.Log, error) {
	getId := FromDomainLog(log)

	trx := rep.Conn.Begin()
	var visitor Visitors
	visitor.PatientsId = getId.PatientsId
	visitor.SchedulesId = getId.SchedulesId

	result := trx.Find(&visitor)
	if result.Error != nil {
		return visitors.Log{}, result.Error
	}

	convertToLog := FromVisitorToLog(visitor)
	convertToLog.Solusi = getId.Solusi
	tambahLog := trx.Create(&convertToLog)
	if tambahLog.Error != nil {
		return visitors.Log{}, trx.Error
	}
	trx.SavePoint("sp1")
	resultHapusVisitor := trx.Delete(&visitor)
	if resultHapusVisitor.Error != nil {
		trx.RollbackTo("sp1")
	}
	trx.Commit()
	return convertToLog.ToDomainLog(), nil

}

func (rep *MysqlVisitorRepository) ModificateVisitor(ctx context.Context, domain visitors.Domain) (visitors.Domain, error) {
	var visitor Visitors
	updateData := FromDomain(domain)
	updateData.UpdateAt = time.Now()
	result := rep.Conn.Model(&visitor).Where("schedules_id = ? AND patients_id = ?", updateData.SchedulesId, updateData.PatientsId).Updates(updateData)
	if result.Error != nil {
		return visitors.Domain{}, result.Error
	}
	return updateData.ToDomain(), nil
}

func (rep *MysqlVisitorRepository) ShowVisitor(ctx context.Context, domain visitors.Domain) (visitors.Domain, error) {
	getId := FromDomain(domain)
	result := rep.Conn.Preload("Patient").Find(&getId)
	if result.Error != nil {
		return visitors.Domain{}, result.Error
	}
	return getId.ToDomain(), nil
}
func (rep *MysqlVisitorRepository) CancelVisitor(ctx context.Context, domain visitors.Domain) (visitors.Domain, error) {
	getId := FromDomain(domain)
	result := rep.Conn.Delete(&getId)
	if result.Error != nil {
		return visitors.Domain{}, result.Error
	}
	return getId.ToDomain(), nil
}
func (rep *MysqlVisitorRepository) DontCome(ctx context.Context, domain visitors.Log) (visitors.Log, error) {
	getId := FromDomainLog(domain)

	trx := rep.Conn.Begin()
	var visitor Visitors
	visitor.PatientsId = getId.PatientsId
	visitor.SchedulesId = getId.SchedulesId

	result := trx.Find(&visitor)
	if result.RowsAffected == 0 {
		return visitors.Log{}, result.Error
	}
	if result.Error != nil {
		return visitors.Log{}, result.Error
	}

	convertToLog := FromVisitorToLog(visitor)
	convertToLog.Message = getId.Message
	tambahLog := trx.Create(&convertToLog)
	if tambahLog.Error != nil {
		return visitors.Log{}, trx.Error
	}
	trx.SavePoint("sp1")
	resultHapusVisitor := trx.Delete(&visitor)
	if resultHapusVisitor.Error != nil {
		trx.RollbackTo("sp1")
	}
	trx.Commit()
	return convertToLog.ToDomainLog(), nil
}

func (rep *MysqlVisitorRepository) ShowAllPatient(ctx context.Context, domain visitors.Domain) ([]visitors.Domain, error) {
	visitorData := []Visitors{}
	visitor := FromDomain(domain)
	result := rep.Conn.Preload("Patient").Where("schedules_id = ?", visitor.SchedulesId).Order("antrian_id asc").Find(&visitorData)
	if result.Error != nil {
		return []visitors.Domain{}, result.Error
	}
	data := ToDomainList(visitorData)
	return data, nil
}
