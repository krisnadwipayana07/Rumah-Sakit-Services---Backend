package patients

import (
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
	BirthDate     time.Time
	BirthPlace    string
	NoBPJS        string
	Token         string
	ContactPerson string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
}
