package patients

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

func (uc *DoctorUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	if domain.Name == "" {
		return Domain{}, errors.New("name empty")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("address empty")
	}
	if domain.BirthDate.String() == "" {
		return Domain{}, errors.New("birth Date empty")
	}
	if domain.BirthPlace == "" {
		return Domain{}, errors.New("birth place empty")
	}
	if domain.NoBPJS == "" {
		return Domain{}, errors.New("no BPJS empty")
	}
	if domain.ContactPerson == "" {
		return Domain{}, errors.New("contact Person empty")
	}

	user, err := uc.Repo.Register(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *DoctorUsecase) Update(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Id == 0 {
		return Domain{}, errors.New("id empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	if domain.Name == "" {
		return Domain{}, errors.New("name empty")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("address empty")
	}
	if domain.BirthDate.String() == "" {
		return Domain{}, errors.New("birthDate empty")
	}
	if domain.BirthPlace == "" {
		return Domain{}, errors.New("birthplace empty")
	}
	if domain.NoBPJS == "" {
		return Domain{}, errors.New("no BPJS empty")
	}
	if domain.ContactPerson == "" {
		return Domain{}, errors.New("contact Person empty")
	}

	user, err := uc.Repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *DoctorUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := uc.Repo.Login(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
