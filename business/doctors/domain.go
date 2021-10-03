package doctors

import (
	"backend/business/schedules"
	"context"
	"database/sql"
	"time"
)

type Domain struct {
	ID            uint
	Email         string
	Password      string
	Name          string
	Address       string
	Nip           string
	Description   string
	DoctorJob     string
	Token         string
	ContactPerson string
	Schedules     []schedules.Domain
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	ShowAll(ctx context.Context) ([]Domain, error)
	// AddSchedule(ctx context.Context, doctorId, scheduleId uint) ([]uint, error)
}

type Repository interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	ShowAll(ctx context.Context) ([]Domain, error)
	// AddSchedule(ctx context.Context, domain Domain, doctorId, scheduleId uint) ([]uint, error)
}
