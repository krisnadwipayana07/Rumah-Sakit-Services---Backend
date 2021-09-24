package patients

import (
	"context"
	"time"
)

type Domain struct {
	Id            int
	Email         string
	Password      string
	Name          string
	Address       string
	BirthDate     time.Time
	BirthPlace    string
	NoBPJS        string
	Token         string
	ContactPerson string
	CreateAt      time.Time
	UpdateAt      time.Time
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
