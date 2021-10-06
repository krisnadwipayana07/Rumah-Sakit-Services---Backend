package visitors_test

import (
	"backend/business"
	"backend/business/visitors"
	"backend/business/visitors/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var visitorRepository mocks.Repository
var visitorServices visitors.Usecase
var visitorDomain visitors.Domain

var visitorLog visitors.Log

func setup() {
	visitorServices = visitors.NewVisitorUsecase(&visitorRepository, time.Hour*1)
	visitorDomain = visitors.Domain{
		SchedulesId: 1,
		PatientsId:  1,
		AntrianId:   1,
		Keluhan:     "Sakit Puchinggg",
	}
	visitorLog = visitors.Log{
		SchedulesId: 1,
		PatientsId:  1,
		AntrianId:   1,
		Keluhan:     "Sakit Puchinggg",
		Solusi:      "Goblok sihh",
	}
}

func TestAddVisitor(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Add Visitor - success", func(t *testing.T) {
		visitorRepository.On("AddVisitor", mock.Anything, mock.Anything).Return(visitorDomain, nil).Once()
		visitor, err := visitorServices.AddVisitor(context.Background(), visitorDomain)
		assert.Nil(t, err)
		assert.Equal(t, visitor.AntrianId, visitorDomain.AntrianId)
	})
	t.Run("Test Case 2 | Add Visitor Failed - domain empty", func(t *testing.T) {
		visitorRepository.On("AddVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.AddVisitor(context.Background(), visitors.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 3 | Add Visitor Failed - patients id empty", func(t *testing.T) {
		visitorRepository.On("AddVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.AddVisitor(context.Background(), visitors.Domain{
			SchedulesId: visitorDomain.SchedulesId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 4 | Add Visitor Failed - keluhan empty", func(t *testing.T) {
		visitorRepository.On("AddVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.AddVisitor(context.Background(), visitors.Domain{
			SchedulesId: visitorDomain.SchedulesId,
			PatientsId:  visitorDomain.PatientsId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 5 | Add Visitor Failed - data empty", func(t *testing.T) {
		visitorRepository.On("AddVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		_, err := visitorServices.AddVisitor(context.Background(), visitorDomain)
		assert.NotNil(t, err)
	})
}

func TestRemoveVisitorToLog(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Remove Visitor - success", func(t *testing.T) {
		visitorRepository.On("RemoveVisitorToLog", mock.Anything, mock.Anything).Return(visitorLog, nil).Once()
		visitor, err := visitorServices.RemoveVisitorToLog(context.Background(), visitorLog)
		assert.Nil(t, err)
		assert.Equal(t, visitor.AntrianId, visitorDomain.AntrianId)
	})
	t.Run("Test Case 2 | Remove Visitor Failed - domain empty", func(t *testing.T) {
		visitorRepository.On("RemoveVisitorToLog", mock.Anything, mock.Anything).Return(visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.RemoveVisitorToLog(context.Background(), visitors.Log{})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Log{})
	})
	t.Run("Test Case 3 | Remove Visitor Failed - patient id empty", func(t *testing.T) {
		visitorRepository.On("RemoveVisitorToLog", mock.Anything, mock.Anything).Return(visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.RemoveVisitorToLog(context.Background(), visitors.Log{
			SchedulesId: visitorLog.SchedulesId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Log{})
	})
	t.Run("Test Case 4 | Remove Visitor Failed - solusi and message empty", func(t *testing.T) {
		visitorRepository.On("RemoveVisitorToLog", mock.Anything, mock.Anything).Return(visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.RemoveVisitorToLog(context.Background(), visitors.Log{
			SchedulesId: visitorLog.SchedulesId,
			PatientsId:  visitorLog.PatientsId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Log{})
	})
	t.Run("Test Case 5 | Remove Visitor Failed - data empty", func(t *testing.T) {
		visitorRepository.On("RemoveVisitorToLog", mock.Anything, mock.Anything).Return(visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.RemoveVisitorToLog(context.Background(), visitors.Log{
			SchedulesId: visitorLog.SchedulesId,
			PatientsId:  visitorLog.PatientsId,
			Solusi:      visitorLog.Solusi,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Log{})
	})
}

func TestModificateVisitor(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Modificate Visitor - success", func(t *testing.T) {
		visitorRepository.On("ModificateVisitor", mock.Anything, mock.Anything).Return(visitorDomain, nil).Once()
		visitor, err := visitorServices.ModificateVisitor(context.Background(), visitorDomain)
		assert.Nil(t, err)
		assert.Equal(t, visitor, visitorDomain)
	})
	t.Run("Test Case 2 | Modificate Visitor Failed - domain empty", func(t *testing.T) {
		visitorRepository.On("ModificateVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ModificateVisitor(context.Background(), visitors.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 3 | Modificate Visitor Failed - schedule id empty", func(t *testing.T) {
		visitorRepository.On("ModificateVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ModificateVisitor(context.Background(), visitors.Domain{
			AntrianId: visitorDomain.AntrianId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 4 | Modificate Visitor Failed - patient id empty", func(t *testing.T) {
		visitorRepository.On("ModificateVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ModificateVisitor(context.Background(), visitors.Domain{
			AntrianId:   visitorDomain.AntrianId,
			SchedulesId: visitorDomain.SchedulesId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 5 | Modificate Visitor Failed - keluhan empty", func(t *testing.T) {
		visitorRepository.On("ModificateVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ModificateVisitor(context.Background(), visitors.Domain{
			AntrianId:   visitorDomain.AntrianId,
			SchedulesId: visitorDomain.SchedulesId,
			PatientsId:  visitorDomain.PatientsId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 5 | Modificate Visitor Failed - keluhan empty", func(t *testing.T) {
		visitorRepository.On("ModificateVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		_, err := visitorServices.ModificateVisitor(context.Background(), visitorDomain)
		assert.NotNil(t, err)
	})
}

func TestShowVisitor(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Show Visitor - success", func(t *testing.T) {
		visitorRepository.On("ShowVisitor", mock.Anything, mock.Anything).Return(visitorDomain, nil).Once()
		visitor, err := visitorServices.ShowVisitor(context.Background(), visitorDomain)
		assert.Nil(t, err)
		assert.Equal(t, visitor, visitorDomain)
	})
	t.Run("Test Case 2 | Show Visitor  failed - domain empty", func(t *testing.T) {
		visitorRepository.On("ShowVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ShowVisitor(context.Background(), visitors.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 3 | Show Visitor  failed - patient id empty", func(t *testing.T) {
		visitorRepository.On("ShowVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ShowVisitor(context.Background(), visitors.Domain{
			SchedulesId: visitorDomain.SchedulesId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 4 | Show Visitor  failed - data empty", func(t *testing.T) {
		visitorRepository.On("ShowVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ShowVisitor(context.Background(), visitorDomain)
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
}

func TestCancelVisitor(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Cancel Visitor  - success", func(t *testing.T) {
		visitorRepository.On("CancelVisitor", mock.Anything, mock.Anything).Return(visitorDomain, nil).Once()
		visitor, err := visitorServices.CancelVisitor(context.Background(), visitorDomain)
		assert.Nil(t, err)
		assert.Equal(t, visitor, visitorDomain)
	})
	t.Run("Test Case 2 | Cancel Visitor  failed - domain empty", func(t *testing.T) {
		visitorRepository.On("CancelVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.CancelVisitor(context.Background(), visitors.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 3 | Cancel Visitor  failed - patient id empty", func(t *testing.T) {
		visitorRepository.On("CancelVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.CancelVisitor(context.Background(), visitors.Domain{
			SchedulesId: visitorDomain.SchedulesId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
	t.Run("Test Case 4 | Cancel Visitor Visitor failed - data empty", func(t *testing.T) {
		visitorRepository.On("CancelVisitor", mock.Anything, mock.Anything).Return(visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.CancelVisitor(context.Background(), visitorDomain)
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Domain{})
	})
}

func TestDontCome(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Dont Come Visitor - success", func(t *testing.T) {
		visitorRepository.On("DontCome", mock.Anything, mock.Anything).Return(visitorLog, nil).Once()
		visitor, err := visitorServices.DontCome(context.Background(), visitorLog)
		assert.Nil(t, err)
		assert.Equal(t, visitorLog, visitor)
	})
	t.Run("Test Case 2 | Dont Come Visitor failed - domain empty", func(t *testing.T) {
		visitorRepository.On("DontCome", mock.Anything, mock.Anything).Return(visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.DontCome(context.Background(), visitors.Log{})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Log{})
	})
	t.Run("Test Case 3 | Dont Come Visitor failed - patient id empty", func(t *testing.T) {
		visitorRepository.On("DontCome", mock.Anything, mock.Anything).Return(visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.DontCome(context.Background(), visitors.Log{
			SchedulesId: visitorLog.SchedulesId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Log{})
	})
	t.Run("Test Case 4 | Dont Come Visitor failed - data empty", func(t *testing.T) {
		visitorRepository.On("DontCome", mock.Anything, mock.Anything).Return(visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.DontCome(context.Background(), visitors.Log{
			SchedulesId: visitorLog.SchedulesId,
			PatientsId:  visitorLog.PatientsId,
		})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, visitors.Log{})
	})
}

func TestShowAllPatient(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Show All Patient  - success", func(t *testing.T) {
		visitorRepository.On("ShowAllPatient", mock.Anything, mock.Anything).Return([]visitors.Domain{visitorDomain}, nil).Once()
		visitor, err := visitorServices.ShowAllPatient(context.Background(), visitorDomain)
		assert.Nil(t, err)
		assert.Equal(t, visitor, []visitors.Domain{visitorDomain})
	})
	t.Run("Test Case 2 | Show All Patient  failed - domain empty", func(t *testing.T) {
		visitorRepository.On("ShowAllPatient", mock.Anything, mock.Anything).Return([]visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ShowAllPatient(context.Background(), visitors.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, []visitors.Domain{})
	})
	t.Run("Test Case 3 | Show All Patient Visitor failed - data empty", func(t *testing.T) {
		visitorRepository.On("ShowAllPatient", mock.Anything, mock.Anything).Return([]visitors.Domain{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ShowAllPatient(context.Background(), visitorDomain)
		assert.NotNil(t, err)
		assert.Equal(t, visitor, []visitors.Domain{})
	})
}

func TestShowLogOfPatient(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Show Log Of Patient  - success", func(t *testing.T) {
		visitorRepository.On("ShowLogOfPatient", mock.Anything, mock.Anything).Return([]visitors.Log{visitorLog}, nil).Once()
		visitor, err := visitorServices.ShowLogOfPatient(context.Background(), visitorLog)
		assert.Nil(t, err)
		assert.Equal(t, visitor, []visitors.Log{visitorLog})
	})
	t.Run("Test Case 2 | Show Log Of Patient  failed - domain empty", func(t *testing.T) {
		visitorRepository.On("ShowLogOfPatient", mock.Anything, mock.Anything).Return([]visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ShowLogOfPatient(context.Background(), visitors.Log{})
		assert.NotNil(t, err)
		assert.Equal(t, visitor, []visitors.Log{})
	})
	t.Run("Test Case 3 | Show Log Of Patient Visitor failed - data empty", func(t *testing.T) {
		visitorRepository.On("ShowLogOfPatient", mock.Anything, mock.Anything).Return([]visitors.Log{}, business.ErrTesting).Once()
		visitor, err := visitorServices.ShowLogOfPatient(context.Background(), visitorLog)
		assert.NotNil(t, err)
		assert.Equal(t, visitor, []visitors.Log{})
	})
}

func TestGetDetailSchedule(t *testing.T) {
	setup()
	last := uint(1)
	count := uint(1)
	t.Run("Test Case 1 | Get Detail Schedule  - success", func(t *testing.T) {
		visitorRepository.On("GetDetailSchedule", mock.Anything, mock.Anything).Return(last, count, nil).Once()
		visitorLast, visitorCount, err := visitorServices.GetDetailSchedule(context.Background(), visitorDomain)
		assert.Nil(t, err)
		assert.Equal(t, visitorLast, last)
		assert.Equal(t, visitorCount, count)
	})
	t.Run("Test Case 2 | Get Detail Schedule  failed - domain empty", func(t *testing.T) {
		visitorRepository.On("GetDetailSchedule", mock.Anything, mock.Anything).Return(0, 0, business.ErrTesting).Once()
		visitorLast, visitorCount, err := visitorServices.GetDetailSchedule(context.Background(), visitors.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, visitorLast, uint(0))
		assert.Equal(t, visitorCount, uint(0))
	})
	//tanyakan?
	// t.Run("Test Case 3 | Get Detail Schedule failed - data empty", func(t *testing.T) {
	// 	visitorRepository.On("GetDetailSchedule", mock.Anything, mock.Anything).Return(0, 0, business.ErrTesting).Once()
	// 	visitorLast, visitorCount, err := visitorServices.GetDetailSchedule(context.Background(), visitorDomain)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, uint(0), visitorLast)
	// 	assert.Equal(t, uint(0), visitorCount)
	// })
}
