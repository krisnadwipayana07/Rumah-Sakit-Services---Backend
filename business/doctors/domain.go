package doctors

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
	Nip           string
	Description   string
	DoctorJob     string
	Token         string
	ContactPerson string
	CreateAt      time.Time
	UpdateAt      time.Time
}

type Usecase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, email, password, name, nip, address, description, doctorJob, contactPerson string) (Domain, error)
	Update(ctx context.Context, id int, name, address, nip, doctorJob, email, token, description, contactPerson string) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, email, password, name, nip, address, description, doctorJob, contactPerson string) (Domain, error)
	Update(ctx context.Context, id int, name, address, nip, doctorJob, email, description, contactPerson string) (Domain, error)
}
