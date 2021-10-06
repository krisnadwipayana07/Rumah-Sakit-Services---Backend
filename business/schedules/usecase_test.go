package schedules_test

import (
	"backend/business"
	"backend/business/patients"
	"backend/business/schedules"
	"backend/business/schedules/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var scheduleRepository mocks.Repository
var scheduleService schedules.Usecase
var scheduleDomain schedules.Domain

func setup() {
	scheduleService = schedules.NewSquedulesUsecase(&scheduleRepository, time.Hour*1)
	scheduleDomain = schedules.Domain{
		ID:          1,
		DoctorId:    1,
		Room:        "lantai 2 Pediatri",
		Message:     "Lagi break makan siang",
		TanggalJaga: time.Now(),
		JamAwal:     "10:00",
		JamAkhir:    "14:00",
		Patients: patients.Domain{
			ID:            1,
			Email:         "krisnadwipayana@gmail.com",
			Password:      "12345678",
			Name:          "Krisna",
			Address:       "gianyar bali",
			BirthDate:     time.Now(),
			BirthPlace:    "Gianyar",
			NoBPJS:        "085641323584321",
			ContactPerson: "087861636910",
			CreatedAt:     time.Now(),
		},
		CreatedAt: time.Date(2021, time.October, 4, 9, 0, 0, 0, time.Local),
		UpdatedAt: time.Now(),
	}
}

func TestAdd(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Add", func(t *testing.T) {
		scheduleRepository.On("Add", mock.Anything, mock.Anything).Return(scheduleDomain, nil).Once()
		schedule, err := scheduleService.Add(context.Background(), scheduleDomain)
		assert.Nil(t, err)
		assert.Equal(t, schedule.ID, scheduleDomain.ID)
	})
	t.Run("Test Case 2 | Failed Add - empty domain", func(t *testing.T) {
		scheduleRepository.On("Add", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Add(context.Background(), schedules.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 3 | Failed Add - DoctorId empty", func(t *testing.T) {
		scheduleRepository.On("Add", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Add(context.Background(), schedules.Domain{
			ID: scheduleDomain.ID,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 4 | Failed Add - Room empty", func(t *testing.T) {
		scheduleRepository.On("Add", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Add(context.Background(), schedules.Domain{
			ID:       scheduleDomain.ID,
			DoctorId: scheduleDomain.DoctorId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 5 | Failed Add - Tanggal Jaga empty", func(t *testing.T) {
		scheduleRepository.On("Add", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Add(context.Background(), schedules.Domain{
			ID:       scheduleDomain.ID,
			DoctorId: scheduleDomain.DoctorId,
			Room:     scheduleDomain.Room,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 6 | Failed Add - Jam Awal empty", func(t *testing.T) {
		scheduleRepository.On("Add", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Add(context.Background(), schedules.Domain{
			ID:          scheduleDomain.ID,
			DoctorId:    scheduleDomain.DoctorId,
			Room:        scheduleDomain.Room,
			TanggalJaga: scheduleDomain.TanggalJaga,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 7 | Failed Add - Jam Akhir empty", func(t *testing.T) {
		scheduleRepository.On("Add", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Add(context.Background(), schedules.Domain{
			ID:          scheduleDomain.ID,
			DoctorId:    scheduleDomain.DoctorId,
			Room:        scheduleDomain.Room,
			TanggalJaga: scheduleDomain.TanggalJaga,
			JamAwal:     scheduleDomain.JamAwal,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 7 | Failed Add - Data empty", func(t *testing.T) {
		scheduleRepository.On("Add", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Add(context.Background(), scheduleDomain)
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
}

func TestRemove(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Remove", func(t *testing.T) {
		scheduleRepository.On("Remove", mock.Anything, mock.Anything).Return(scheduleDomain, nil).Once()
		schedule, err := scheduleService.Remove(context.Background(), scheduleDomain)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), schedule.ID)
	})
	t.Run("Test Case 2 | Remove Failed - id Empty", func(t *testing.T) {
		scheduleRepository.On("Remove", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Remove(context.Background(), schedules.Domain{
			ID: 0,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 3 | Remove Failed - Data Empty", func(t *testing.T) {
		scheduleRepository.On("Remove", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Remove(context.Background(), scheduleDomain)
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
}

func TestModificate(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Modificate", func(t *testing.T) {
		scheduleRepository.On("Modificate", mock.Anything, mock.Anything).Return(scheduleDomain, nil).Once()
		schedule, err := scheduleService.Modificate(context.Background(), scheduleDomain)
		assert.Nil(t, err)
		assert.Equal(t, schedule.ID, scheduleDomain.ID)
	})
	t.Run("Test Case 2 | Failed Modificate - empty domain", func(t *testing.T) {
		scheduleRepository.On("Modificate", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Modificate(context.Background(), schedules.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 2 | Failed Modificate - DoctorID Empty", func(t *testing.T) {
		scheduleRepository.On("Modificate", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Modificate(context.Background(), schedules.Domain{
			ID: scheduleDomain.ID,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 2 | Failed Modificate - Room Empty", func(t *testing.T) {
		scheduleRepository.On("Modificate", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Modificate(context.Background(), schedules.Domain{
			ID:       scheduleDomain.ID,
			DoctorId: scheduleDomain.DoctorId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 3 | Failed Modificate - Tanggal Jaga empty", func(t *testing.T) {
		scheduleRepository.On("Modificate", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Modificate(context.Background(), schedules.Domain{
			ID:       scheduleDomain.ID,
			DoctorId: scheduleDomain.DoctorId,
			Room:     scheduleDomain.Room,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 4 | Failed Modificate - Jam Awal empty", func(t *testing.T) {
		scheduleRepository.On("Modificate", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Modificate(context.Background(), schedules.Domain{
			ID:          scheduleDomain.ID,
			DoctorId:    scheduleDomain.DoctorId,
			Room:        scheduleDomain.Room,
			TanggalJaga: scheduleDomain.TanggalJaga,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 5 | Failed Modificate - Jam Akhir empty", func(t *testing.T) {
		scheduleRepository.On("Modificate", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Modificate(context.Background(), schedules.Domain{
			ID:          scheduleDomain.ID,
			DoctorId:    scheduleDomain.DoctorId,
			Room:        scheduleDomain.Room,
			TanggalJaga: scheduleDomain.TanggalJaga,
			JamAwal:     scheduleDomain.JamAwal,
		})
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
	t.Run("Test Case 5 | Failed Modificate - data empty", func(t *testing.T) {
		scheduleRepository.On("Modificate", mock.Anything, mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		schedule, err := scheduleService.Modificate(context.Background(), scheduleDomain)
		assert.NotNil(t, err)
		assert.Equal(t, schedule, schedules.Domain{})
	})
}

func TestShow(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Show Data", func(t *testing.T) {
		scheduleRepository.On("Show",
			mock.Anything,
			mock.Anything).Return(scheduleDomain, nil).Once()
		schedule, err := scheduleService.Show(context.Background(), scheduleDomain)
		assert.Nil(t, err)
		assert.Equal(t, schedule, scheduleDomain)
	})
	t.Run("Test Case 2 | Failed Show - empty domain", func(t *testing.T) {
		scheduleRepository.On("Show",
			mock.Anything,
			mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		_, err := scheduleService.Show(context.Background(), schedules.Domain{})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 2 | Failed Show - empty data", func(t *testing.T) {
		scheduleRepository.On("Show", mock.Anything,
			mock.Anything).Return(schedules.Domain{}, business.ErrTesting).Once()
		_, err := scheduleService.Show(context.Background(), scheduleDomain)
		assert.NotNil(t, err)
	})
}

func TestGetAllInDoctor(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Get All Doctor - Success", func(t *testing.T) {
		scheduleRepository.On("GetAllInOneDoctor", mock.Anything, mock.Anything).Return(
			[]schedules.Domain{scheduleDomain}, nil).Once()
		schedule, err := scheduleService.GetAllInOneDoctor(context.Background(), scheduleDomain)
		assert.Nil(t, err)
		assert.Equal(t, schedule, []schedules.Domain{scheduleDomain})
	})
	t.Run("Test Case 2 | Get All Doctor Failed - domain Empty", func(t *testing.T) {
		scheduleRepository.On("GetAllInOneDoctor", mock.Anything, mock.Anything).Return(
			[]schedules.Domain{}, business.ErrTesting).Once()
		_, err := scheduleService.GetAllInOneDoctor(context.Background(), schedules.Domain{
			ID: 0,
		})
		assert.NotNil(t, err)

	})
	t.Run("Test Case 3 | Get All Doctor Failed - Data Empty", func(t *testing.T) {
		scheduleRepository.On("GetAllInOneDoctor", mock.Anything, schedules.Domain{}).Return(
			[]schedules.Domain{}, business.ErrTesting).Once()
		_, err := scheduleService.GetAllInOneDoctor(context.Background(), scheduleDomain)
		assert.NotNil(t, err)
	})
	// t.Run("Test Case 3 | Get All Doctor Failed - Data Empty", func(t *testing.T) {
	// 	scheduleRepository.On("GetAllInOneDoctor", mock.Anything, schedules.Domain{}).Return(
	// 		[]schedules.Domain{}, business.ErrTesting).Once()
	// 	schedule, err := scheduleService.GetAllInOneDoctor(context.Background(), scheduleDomain)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, scheduleDomain, schedule)
	// })
}

func TestGetAll(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Get All Doctor - Success", func(t *testing.T) {
		scheduleRepository.On("GetAll", mock.Anything, mock.Anything).Return(
			[]schedules.Domain{scheduleDomain}, nil).Once()
		schedule, err := scheduleService.GetAll(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, schedule, []schedules.Domain{scheduleDomain})
	})
	t.Run("Test Case 2 | Get All Doctor Failed - domain Empty", func(t *testing.T) {
		scheduleRepository.On("GetAll", mock.Anything, mock.Anything).Return(
			[]schedules.Domain{}, business.ErrTesting).Once()
		_, err := scheduleService.GetAll(context.Background())
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Get All Doctor Failed - Data Empty", func(t *testing.T) {
		scheduleRepository.On("GetAll", mock.Anything, mock.Anything).Return(
			[]schedules.Domain{}, business.ErrTesting).Once()
		_, err := scheduleService.GetAll(context.Background())
		assert.NotNil(t, err)
	})
}
