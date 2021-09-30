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
	// if result.RowsAffected == 0 {
	// 	newVisitor.AntrianId = 1
	// 	fmt.Println("ga work")
	// } else {
	// 	newVisitor.AntrianId = newVisitor.AntrianId + 1
	// 	fmt.Println("work")
	// }
	fmt.Println(newVisitor.AntrianId)
	// newVisitor.AntrianId = 1
	newVisitor.CreateAt = time.Now()
	resultVisitor := rep.Conn.Create(&newVisitor)
	if resultVisitor.Error != nil {
		return visitors.Domain{}, resultVisitor.Error
	}
	return newVisitor.ToDomain(), nil

}
