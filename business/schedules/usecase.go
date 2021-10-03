package schedules

import (
	"context"
	"errors"
	"time"
)

type SquedulesUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewSquedulesUsecase(repo Repository, timeout time.Duration) Usecase {
	return &SquedulesUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *SquedulesUsecase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.DoctorId <= 0 {
		return Domain{}, errors.New("doctorId empty")
	}
	if domain.Room == "" {
		return Domain{}, errors.New("room empty")
	}
	if domain.TanggalJaga.String() == "" {
		return Domain{}, errors.New("tanggal jaga empty")
	}
	if domain.JamAwal == "" {
		return Domain{}, errors.New("jam jaga awal empty")
	}
	if domain.JamAkhir == "" {
		return Domain{}, errors.New("jam jaga akhir empty")
	}

	squedule, err := uc.Repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return squedule, nil
}

func (uc *SquedulesUsecase) Remove(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID <= 0 {
		return Domain{}, errors.New("id empty")
	}

	squedules, err := uc.Repo.Remove(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return squedules, nil
}

func (uc *SquedulesUsecase) Modificate(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID <= 0 {
		return Domain{}, errors.New("id empty")
	}
	if domain.DoctorId <= 0 {
		return Domain{}, errors.New("doctorId empty")
	}
	if domain.Room == "" {
		return Domain{}, errors.New("room empty")
	}
	if domain.TanggalJaga.String() == "" {
		return Domain{}, errors.New("tanggal jaga empty")
	}
	if domain.JamAwal == "" {
		return Domain{}, errors.New("jam jaga awal empty")
	}
	if domain.JamAkhir == "" {
		return Domain{}, errors.New("jam jaga akhir empty")
	}

	squedule, err := uc.Repo.Modificate(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return squedule, nil
}

func (uc *SquedulesUsecase) Show(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID <= 0 {
		return Domain{}, errors.New("id empty")
	}

	squedule, err := uc.Repo.Show(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return squedule, nil
}

func (uc *SquedulesUsecase) GetAllInOneDoctor(ctx context.Context, domain Domain) ([]Domain, error) {
	if domain.DoctorId <= 0 {
		return []Domain{}, errors.New("doctorId empty")
	}
	allData, err := uc.Repo.GetAllInOneDoctor(ctx, domain)
	if err != nil {
		return []Domain{}, err
	}
	return allData, nil
}
func (uc *SquedulesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	allData, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return allData, nil
}

// func (uc *SquedulesUsecase) InsertDoctor(ctx context.Context, domain Domain) (Domain, error) {
// 	if domain.DoctorId <= 0 {
// 		return Domain{}, errors.New("doctor id empty")
// 	}
// 	if domain.ID <= 0 {
// 		return Domain{}, errors.New("id empty")
// 	}
// 	schedule, err := uc.Repo.InsertDoctor(ctx, domain)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	return schedule, nil
// }
