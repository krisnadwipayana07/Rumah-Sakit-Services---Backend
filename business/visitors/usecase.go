package visitors

import (
	"context"
	"errors"
	"time"
)

type VisitorUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewVisitorUsecase(repo Repository, timeout time.Duration) Usecase {
	return &VisitorUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (uc *VisitorUsecase) AddVisitor(ctx context.Context, domain Domain) (Domain, error) {
	// if domain.AntrianId <= 0 {
	// 	return Domain{}, errors.New("antrean Id empty")
	// }
	if domain.SchedulesId <= 0 {
		return Domain{}, errors.New("schedule Id empty")
	}
	if domain.PatientsId <= 0 {
		return Domain{}, errors.New("patient Id empty")
	}
	if domain.Keluhan == "" {
		return Domain{}, errors.New("keluhan empty")
	}

	visitor, err := uc.Repo.AddVisitor(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return visitor, nil
}
func (uc *VisitorUsecase) RemoveVisitorToLog(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID <= 0 {
		return Domain{}, errors.New("id empty")
	}
	visitor, err := uc.Repo.AddVisitor(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return visitor, nil
}
func (uc *VisitorUsecase) ModificateVisitor(ctx context.Context, domain Domain) (Domain, error) {

	// if domain.ID <= 0 {
	// 	return Domain{}, errors.New("id empty")
	// }
	if domain.AntrianId <= 0 {
		return Domain{}, errors.New("antrean Id empty")
	}
	if domain.SchedulesId <= 0 {
		return Domain{}, errors.New("schedule Id empty")
	}
	if domain.PatientsId <= 0 {
		return Domain{}, errors.New("patient Id empty")
	}
	if domain.Keluhan == "" {
		return Domain{}, errors.New("keluhan empty")
	}

	visitor, err := uc.Repo.AddVisitor(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return visitor, nil
}
func (uc *VisitorUsecase) ShowVisitor(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID <= 0 {
		return Domain{}, errors.New("id empty")
	}
	visitor, err := uc.Repo.AddVisitor(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return visitor, nil
}
