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
func (uc *VisitorUsecase) RemoveVisitorToLog(ctx context.Context, log Log) (Log, error) {
	if log.SchedulesId <= 0 {
		return Log{}, errors.New("schedule id empty")
	}
	if log.PatientsId <= 0 {
		return Log{}, errors.New("patient id empty")
	}
	visitor, err := uc.Repo.RemoveVisitorToLog(ctx, log)
	if err != nil {
		return Log{}, err
	}
	return visitor, nil
}
func (uc *VisitorUsecase) ModificateVisitor(ctx context.Context, domain Domain) (Domain, error) {
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

	visitor, err := uc.Repo.ModificateVisitor(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return visitor, nil
}

func (uc *VisitorUsecase) ShowVisitor(ctx context.Context, domain Domain) (Domain, error) {
	if domain.SchedulesId <= 0 {
		return Domain{}, errors.New("schedule id empty")
	}
	if domain.PatientsId <= 0 {
		return Domain{}, errors.New("patient id empty")
	}
	visitor, err := uc.Repo.ShowVisitor(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return visitor, nil
}

func (uc *VisitorUsecase) CancelVisitor(ctx context.Context, domain Domain) (Domain, error) {
	if domain.SchedulesId <= 0 {
		return Domain{}, errors.New("schedule id empty")
	}
	if domain.PatientsId <= 0 {
		return Domain{}, errors.New("patient id empty")
	}
	data, err := uc.Repo.CancelVisitor(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}
func (uc *VisitorUsecase) DontCome(ctx context.Context, log Log) (Log, error) {
	if log.SchedulesId <= 0 {
		return Log{}, errors.New("schedule id empty")
	}
	if log.PatientsId <= 0 {
		return Log{}, errors.New("patient id empty")
	}
	data, err := uc.Repo.DontCome(ctx, log)
	if err != nil {
		return Log{}, err
	}
	return data, nil
}

func (uc *VisitorUsecase) ShowAllPatient(ctx context.Context, domain Domain) ([]Domain, error) {
	if domain.SchedulesId <= 0 {
		return []Domain{}, errors.New("schedule id empty")
	}
	visitor, err := uc.Repo.ShowAllPatient(ctx, domain)
	if err != nil {
		return []Domain{}, err
	}
	return visitor, nil
}

func (uc *VisitorUsecase) ShowLogOfPatient(ctx context.Context, log Log) ([]Log, error) {
	if log.PatientsId <= 0 {
		return []Log{}, errors.New("patient Id empty")
	}

	data, err := uc.Repo.ShowLogOfPatient(ctx, log)
	if err != nil {
		return []Log{}, err
	}
	return data, nil
}

func (uc *VisitorUsecase) GetDetailSchedule(ctx context.Context, domain Domain) (uint, uint, error) {
	if domain.SchedulesId <= 0 {
		return 0, 0, errors.New("schedule id empty")
	}
	last, count, err := uc.Repo.GetDetailSchedule(ctx, domain)
	if err != nil {
		return 0, 0, err
	}
	return last, count, nil
}
