package patients_test

import (
	"backend/app/middlewares"
	"backend/business"
	"backend/business/patients"
	"backend/business/patients/mocks"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var patientRepository mocks.Repository
var patientService patients.Usecase
var patientDomain patients.Domain

func setup() {
	patientService = patients.NewPatientsUsecase(&patientRepository, &middlewares.ConfigJWT{}, time.Hour*1)
	patientDomain = patients.Domain{
		ID:            1,
		Email:         "krisnadwipayana30@gmail.com",
		Password:      "12345678",
		Name:          "krisna",
		Address:       "Gianyar",
		BirthDate:     time.Date(2001, time.February, 07, 0, 0, 0, 0, time.Local),
		BirthPlace:    "Gianyar",
		NoBPJS:        "087526372323",
		Token:         "test",
		ContactPerson: "038482093490",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     sql.NullTime{},
	}
}

func TestRegister(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Register - Success", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patientDomain, nil).Once()
		patient, err := patientService.Register(context.Background(), patientDomain)
		assert.Nil(t, err)
		assert.Equal(t, patient.Name, patientDomain.Name)
	})
	t.Run("Test Case 2 | Register Failed - domain Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patients.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 3 | Register Failed - Password Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patients.Domain{
			Email: patientDomain.Email,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 4 | Register Failed - Name Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patients.Domain{
			Email:    patientDomain.Email,
			Password: patientDomain.Password,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 5 | Register Failed - Address Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patients.Domain{
			Email:    patientDomain.Email,
			Password: patientDomain.Password,
			Name:     patientDomain.Name,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 6 | Register Failed - Birthday Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patients.Domain{
			Email:    patientDomain.Email,
			Password: patientDomain.Password,
			Name:     patientDomain.Name,
			Address:  patientDomain.Address,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 7 | Register Failed - Birthplace Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patients.Domain{
			Email:     patientDomain.Email,
			Password:  patientDomain.Password,
			Name:      patientDomain.Name,
			Address:   patientDomain.Address,
			BirthDate: patientDomain.BirthDate,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 8 | Register Failed - NoBPJS Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patients.Domain{
			Email:      patientDomain.Email,
			Password:   patientDomain.Password,
			Name:       patientDomain.Name,
			Address:    patientDomain.Address,
			BirthDate:  patientDomain.BirthDate,
			BirthPlace: patientDomain.BirthPlace,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 9 | Register Failed - Contact Person Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patients.Domain{
			Email:      patientDomain.Email,
			Password:   patientDomain.Password,
			Name:       patientDomain.Name,
			Address:    patientDomain.Address,
			BirthDate:  patientDomain.BirthDate,
			BirthPlace: patientDomain.BirthPlace,
			NoBPJS:     patientDomain.NoBPJS,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 9 | Register Failed - data Empty", func(t *testing.T) {
		patientRepository.On("Register", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Register(context.Background(), patientDomain)
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})

}

func TestUpdate(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patientDomain, nil).Once()
		patient, err := patientService.Update(context.Background(), patientDomain)
		assert.Nil(t, err)
		assert.Equal(t, patient.Name, patientDomain.Name)
	})
	t.Run("Test Case 2 | Update Failed - domain Empty", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Update(context.Background(), patients.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 3 | Update Failed - Name Empty", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Update(context.Background(), patients.Domain{
			Email: patientDomain.Email,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 3 | Update Failed - Address Empty", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Update(context.Background(), patients.Domain{
			Email: patientDomain.Email,
			Name:  patientDomain.Name,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 3 | Update Failed - BirthDate Empty", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Update(context.Background(), patients.Domain{
			Email:   patientDomain.Email,
			Name:    patientDomain.Name,
			Address: patientDomain.Address,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 3 | Update Failed - BirthPlace Empty", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Update(context.Background(), patients.Domain{
			Email:     patientDomain.Email,
			Name:      patientDomain.Name,
			Address:   patientDomain.Address,
			BirthDate: patientDomain.BirthDate,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 3 | Update Failed - NoBPJS Empty", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Update(context.Background(), patients.Domain{
			Email:      patientDomain.Email,
			Name:       patientDomain.Name,
			Address:    patientDomain.Address,
			BirthDate:  patientDomain.BirthDate,
			BirthPlace: patientDomain.BirthPlace,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 3 | Update Failed - Contact Person Empty", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Update(context.Background(), patients.Domain{
			Email:      patientDomain.Email,
			Name:       patientDomain.Name,
			Address:    patientDomain.Address,
			BirthDate:  patientDomain.BirthDate,
			BirthPlace: patientDomain.BirthPlace,
			NoBPJS:     patientDomain.NoBPJS,
		})
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
	t.Run("Test Case 4 | Update Failed - data Empty", func(t *testing.T) {
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Update(context.Background(), patientDomain)
		assert.NotNil(t, err)
		assert.Equal(t, patient, patients.Domain{})
	})
}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Login - Success", func(t *testing.T) {
		patientRepository.On("Login", mock.Anything, mock.Anything).Return(patientDomain, nil).Once()
		patient, err := patientService.Login(context.Background(), patients.Domain{
			Email:    patientDomain.Email,
			Password: patientDomain.Password,
		})
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patientDomain, nil).Once()
		patientService.Update(context.Background(), patient)
		assert.Nil(t, err)
		assert.Equal(t, patientDomain.Name, patient.Name)
	})
	t.Run("Test Case 2 | Login Failed - Email Empty", func(t *testing.T) {
		patientRepository.On("Login", mock.Anything, mock.Anything).Return(patientDomain, nil).Once()
		patient, err := patientService.Login(context.Background(), patients.Domain{})
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patientDomain, nil).Once()
		patientService.Update(context.Background(), patient)
		assert.NotNil(t, err)
		assert.Equal(t, patients.Domain{}, patient)
	})
	t.Run("Test Case 2 | Login Failed - Password Empty", func(t *testing.T) {
		patientRepository.On("Login", mock.Anything, mock.Anything).Return(patientDomain, nil).Once()
		patient, err := patientService.Login(context.Background(), patients.Domain{
			Email:    patientDomain.Email,
			Password: "",
		})
		patientRepository.On("Update", mock.Anything, mock.Anything).Return(patientDomain, nil).Once()
		patientService.Update(context.Background(), patient)
		assert.NotNil(t, err)
		assert.Equal(t, patients.Domain{}, patient)
	})
	t.Run("Test Case 3 | Login Failed - Data Empty", func(t *testing.T) {
		patientRepository.On("Login", mock.Anything, mock.Anything).Return(patients.Domain{}, business.ErrTesting).Once()
		patient, err := patientService.Login(context.Background(), patientDomain)
		patientRepository.On("Update", mock.Anything, patient).Return(patients.Domain{}, business.ErrTesting).Once()
		patientService.Update(context.Background(), patient)
		assert.Nil(t, err)
	})
}
