package doctors_test

import (
	"backend/app/middlewares"
	"backend/business"
	"backend/business/doctors"
	"backend/business/doctors/mocks"
	"backend/business/schedules"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var doctorRepository mocks.Repository
var doctorService doctors.Usecase
var doctorDomain doctors.Domain

func setup() {
	doctorService = doctors.NewDoctorUsecase(&doctorRepository, &middlewares.ConfigJWT{}, time.Hour*1)
	doctorDomain = doctors.Domain{
		ID:            1,
		Email:         "krisnadwipayana30@gmail.com",
		Password:      "12345678",
		Name:          "krisna",
		Address:       "Gianyar",
		Nip:           "23456789",
		Description:   "keren dan ganteng",
		DoctorJob:     "backend scientise",
		Token:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicm9sZSI6ImRvY3RvciIsImV4cCI6MTYzMzQzMDU2NX0.2FARkaNYLaHtfnEoeyTaSCw2eTgAwJhltyil9C2l-bY",
		ContactPerson: "0873824342923",
		Schedules:     schedules.Domain{},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     sql.NullTime{},
	}
}

// fungsinya manggil 2 usecase?
func TestLogin(t *testing.T) {
	t.Run("Test case - 1 | Login - Succses", func(t *testing.T) {
		setup()
		doctorRepository.On("Login", mock.Anything, mock.Anything).Return(doctorDomain, nil).Once()
		doctor, err := doctorService.Login(context.Background(), doctors.Domain{
			Email:    "krisnadwipayana30@gmail.com",
			Password: "12345678",
		})
		doctorRepository.On("Update", mock.Anything, doctor).Return(doctorDomain, nil).Once()
		// doctorService.Update(context.Background(), doctorDomain)
		assert.Nil(t, err)
		assert.Equal(t, "krisna", doctor.Name)
	})
	t.Run("Test case - 2 | Login Failed - Email empty", func(t *testing.T) {
		setup()
		doctorRepository.On("Login", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting).Once()
		doctor, err := doctorService.Login(context.Background(), doctors.Domain{
			Email:    "",
			Password: doctorDomain.Password,
		})
		doctorRepository.On("Update", mock.Anything, doctor).Return(doctorDomain, business.ErrTesting).Once()
		// doctorService.Update(context.Background(), doctorDomain)
		assert.NotNil(t, err)
	})
	t.Run("Test case - 3 | Login Failed - Password empty", func(t *testing.T) {
		setup()
		doctorRepository.On("Login", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting).Once()
		doctor, err := doctorService.Login(context.Background(), doctors.Domain{
			Email:    doctorDomain.Email,
			Password: "",
		})
		doctorRepository.On("Update", mock.Anything, doctor).Return(doctorDomain, business.ErrTesting).Once()
		// doctorService.Update(context.Background(), doctorDomain)
		assert.NotNil(t, err)
	})
	t.Run("Test case - 3 | Login Failed - data empty", func(t *testing.T) {
		setup()
		doctorRepository.On("Login", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting).Once()
		doctor, err := doctorService.Login(context.Background(), doctorDomain)
		doctorRepository.On("Update", mock.Anything, doctor).Return(doctors.Domain{}, business.ErrTesting).Once()
		// doctorService.Update(context.Background(), doctorDomain)
		assert.NotNil(t, err)
	})
}

func TestRegister(t *testing.T) {
	setup()
	t.Run("Test case - 1 | Register - Succses", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctorDomain, nil).Once()
		doctor, err := doctorService.Register(context.Background(), doctorDomain)
		assert.Nil(t, err)
		assert.Equal(t, "krisna", doctor.Name)
	})
	t.Run("Test case - 2 | Register - Failed Email empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting).Once()
		_, err := doctorService.Register(context.Background(), doctors.Domain{})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 3 | Register - Failed Password empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting).Once()
		_, err := doctorService.Register(context.Background(), doctors.Domain{
			Email: "krisnadwipayana30@gmail.com",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 4 | Register - Failed Name empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting).Once()
		_, err := doctorService.Register(context.Background(), doctors.Domain{
			Email:    "krisnadwipayana30@gmail.com",
			Password: "12345678",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 5 | Register - Failed Address empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Register(context.Background(), doctors.Domain{
			Email:    "krisnadwipayana30@gmail.com",
			Password: "12345678",
			Name:     "krisna",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 6 | Register - Failed Address empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Register(context.Background(), doctors.Domain{
			Email:    "krisnadwipayana30@gmail.com",
			Password: "12345678",
			Name:     "krisna",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 7 | Register - Failed Nip empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Register(context.Background(), doctors.Domain{
			Email:    "krisnadwipayana30@gmail.com",
			Password: "12345678",
			Name:     "krisna",
			Address:  "Gianyar",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 8 | Register - Failed DoctorJob empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Register(context.Background(), doctors.Domain{
			Email:    "krisnadwipayana30@gmail.com",
			Password: "12345678",
			Name:     "krisna",
			Address:  "Gianyar",
			Nip:      "1232354",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 9 | Register - Failed Contact Person empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Register(context.Background(), doctors.Domain{
			Email:     "krisnadwipayana30@gmail.com",
			Password:  "12345678",
			Name:      "krisna",
			Address:   "Gianyar",
			Nip:       "1232354",
			DoctorJob: "psikolog",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 9 | Register - data empty", func(t *testing.T) {
		doctorRepository.On("Register", mock.Anything, mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		doctor, err := doctorService.Register(context.Background(), doctorDomain)
		assert.NotNil(t, err)
		assert.Equal(t, doctor, doctors.Domain{})
	})
}

func TestUpdate(t *testing.T) {
	setup()
	t.Run("Test case - 1 | Update - Succses", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctorDomain, nil).Once()
		doctor, err := doctorService.Update(context.Background(), doctorDomain)
		assert.Nil(t, err)
		assert.Equal(t, "krisna", doctor.Name)
	})
	t.Run("Test case - 2 | Update - Failed Domain Empty", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Update(context.Background(), doctors.Domain{})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 3 | Update - Failed Name Empty", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Update(context.Background(), doctors.Domain{
			ID: doctorDomain.ID,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 4 | Update - Failed Name Empty", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Update(context.Background(), doctors.Domain{
			ID:   doctorDomain.ID,
			Name: doctorDomain.Name,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 5 | Update - Failed DoctorJob Empty", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Update(context.Background(), doctors.Domain{
			ID:      doctorDomain.ID,
			Name:    doctorDomain.Name,
			Address: doctorDomain.Address,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 6 | Update - Failed Email Empty", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Update(context.Background(), doctors.Domain{
			ID:        doctorDomain.ID,
			Name:      doctorDomain.Name,
			Address:   doctorDomain.Address,
			DoctorJob: doctorDomain.DoctorJob,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 7 | Update - Failed Token Empty", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Update(context.Background(), doctors.Domain{
			ID:        doctorDomain.ID,
			Name:      doctorDomain.Name,
			Address:   doctorDomain.Address,
			DoctorJob: doctorDomain.DoctorJob,
			Email:     doctorDomain.Email,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 8 | Update - Failed Contact Person Empty", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Update(context.Background(), doctors.Domain{
			ID:        doctorDomain.ID,
			Name:      doctorDomain.Name,
			Address:   doctorDomain.Address,
			DoctorJob: doctorDomain.DoctorJob,
			Email:     doctorDomain.Email,
			Token:     doctorDomain.Token,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test case - 8 | Update - Failed Contact Person Empty", func(t *testing.T) {
		doctorRepository.On("Update", mock.Anything,
			mock.Anything).Return(doctors.Domain{}, business.ErrTesting, nil).Once()
		_, err := doctorService.Update(context.Background(), doctors.Domain{
			ID:            doctorDomain.ID,
			Name:          doctorDomain.Name,
			Address:       doctorDomain.Address,
			DoctorJob:     doctorDomain.DoctorJob,
			Email:         doctorDomain.Email,
			Token:         doctorDomain.Token,
			ContactPerson: doctorDomain.ContactPerson,
		})
		assert.NotNil(t, err)
	})
}

func TestShowAll(t *testing.T) {
	setup()
	t.Run("Test case - 1 | Show - Succses", func(t *testing.T) {
		doctorRepository.On("ShowAll",
			mock.Anything).Return([]doctors.Domain{doctorDomain}, nil).Once()
		_, err := doctorService.ShowAll(context.Background())
		assert.Nil(t, err)
	})
	t.Run("Test case - 1 | Show - failed", func(t *testing.T) {
		doctorRepository.On("ShowAll",
			mock.Anything).Return([]doctors.Domain{}, business.ErrTesting).Once()
		_, err := doctorService.ShowAll(context.Background())
		assert.NotNil(t, err)
	})
}
