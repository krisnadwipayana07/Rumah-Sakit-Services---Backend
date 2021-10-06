package doctors

import (
	"backend/app/middlewares"
	"context"
	"errors"
	"time"
)

type DoctorUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewDoctorUsecase(repo Repository, jwtAuth *middlewares.ConfigJWT, timeout time.Duration) Usecase {
	return &DoctorUsecase{
		Repo:           repo,
		contextTimeout: timeout,
		jwtAuth:        jwtAuth,
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
	if domain.Nip == "" {
		return Domain{}, errors.New("nip empty")
	}
	if domain.DoctorJob == "" {
		return Domain{}, errors.New("doctorJob empty")
	}
	if domain.ContactPerson == "" {
		return Domain{}, errors.New("contact person empty")
	}

	user, err := uc.Repo.Register(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *DoctorUsecase) Update(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID <= 0 {
		return Domain{}, errors.New("id Empty")
	}
	if domain.Name == "" {
		return Domain{}, errors.New("name empty")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("address empty")
	}
	if domain.DoctorJob == "" {
		return Domain{}, errors.New("doctorJob empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Token == "" {
		return Domain{}, errors.New("token empty")
	}
	if domain.ContactPerson == "" {
		return Domain{}, errors.New("contact person empty")
	}

	user, err := uc.Repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *DoctorUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

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

	user.Token = uc.jwtAuth.GenerateToken(user.ID, "doctor")

	return user, nil
}

func (uc *DoctorUsecase) ShowAll(ctx context.Context) ([]Domain, error) {
	allData, err := uc.Repo.ShowAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return allData, nil
}

// func (uc *DoctorUsecase) AddSchedule(ctx context.Context, doctorId, scheduleId uint) ([]uint, error) {
// 	if scheduleId <= 0 {
// 		return []uint{0, 0}, errors.New("schedule id empty")
// 	}
// 	if doctorId <= 0 {
// 		return []uint{0, 0}, errors.New("doctor id empty")
// 	}
// 	var domain Domain

// 	manySchedule, err := uc.Repo.AddSchedule(ctx, domain, doctorId, scheduleId)
// 	if err != nil {
// 		return []uint{0, 0}, err
// 	}
// 	return manySchedule, nil
// }
