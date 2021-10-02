package visitors

import (
	"context"
	"time"
)

type Domain struct {
	// ID          uint
	SchedulesId uint
	PatientsId  uint
	AntrianId   uint
	Keluhan     string
	CreateAt    time.Time
	UpdateAt    time.Time
}

type Log struct {
	SchedulesId uint
	PatientsId  uint
	AntrianId   uint
	Keluhan     string
	Solusi      string
	Message     string
	CreateAt    time.Time
	UpdateAt    time.Time
}

type Usecase interface {
	AddVisitor(ctx context.Context, domain Domain) (Domain, error)
	RemoveVisitorToLog(ctx context.Context, log Log) (Log, error)
	ModificateVisitor(ctx context.Context, domain Domain) (Domain, error)
	ShowVisitor(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	AddVisitor(ctx context.Context, domain Domain) (Domain, error)
	RemoveVisitorToLog(ctx context.Context, log Log) (Log, error)
	ModificateVisitor(ctx context.Context, domain Domain) (Domain, error)
	ShowVisitor(ctx context.Context, domain Domain) (Domain, error)
}
