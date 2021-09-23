package doctors

import (
	"context"
	"errors"
	"time"
)

type DoctorUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewDoctorUsecase(repo Repository, timeout time.Duration) Usecase {
	return &DoctorUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *DoctorUsecase) Register(ctx context.Context, email, password, name, nip, address, description, doctorJob, contactPerson string) (Domain, error) {

	if email == "" {
		return Domain{}, errors.New("email empty")
	}
	if password == "" {
		return Domain{}, errors.New("password empty")
	}
	if name == "" {
		return Domain{}, errors.New("name empty")
	}
	if address == "" {
		return Domain{}, errors.New("address empty")
	}
	if doctorJob == "" {
		return Domain{}, errors.New("doctorJob empty")
	}
	if contactPerson == "" {
		return Domain{}, errors.New("contact person empty")
	}

	user, err := uc.Repo.Register(ctx, email, password, name, nip, address, description, doctorJob, contactPerson)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *DoctorUsecase) Update(ctx context.Context, id int, name, address, nip, doctorJob, email, token, description, contactPerson string) (Domain, error) {

	if id == 0 {
		return Domain{}, errors.New("id empty")
	}
	if name == "" {
		return Domain{}, errors.New("name empty")
	}
	if address == "" {
		return Domain{}, errors.New("address empty")
	}
	if doctorJob == "" {
		return Domain{}, errors.New("doctorJob empty")
	}
	if email == "" {
		return Domain{}, errors.New("email empty")
	}
	if token == "" {
		return Domain{}, errors.New("token empty")
	}
	if contactPerson == "" {
		return Domain{}, errors.New("contact person empty")
	}

	user, err := uc.Repo.Update(ctx, id, name, address, nip, doctorJob, email, description, contactPerson)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *DoctorUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email empty")
	}
	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := uc.Repo.Login(ctx, email, password)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
