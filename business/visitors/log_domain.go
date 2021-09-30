package visitors

import (
	"database/sql"
	"time"
)

type VisitorLogs struct {
	ID         uint
	AntrianId  uint
	ScheduleId uint
	DoctorId   uint
	PatientId  uint
	Keluhan    string
	Solusi     string
	CreateAt   time.Time
	UpdateAt   time.Time
	DeleteAt   sql.NullTime
}
