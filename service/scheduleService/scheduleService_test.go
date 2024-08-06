package scheduleservice_test

import (
	"errors"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"go-telemedicine/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateSchedule(t *testing.T) {

	ts := test.SetupTestCase(t)

	req := models.ScheduleCreateRequest{
		DoctorID:  1,
		Date:      "2022-01-01",
		StartTime: "09:00",
		EndTime:   "12:00",
	}

	newData := models.ScheduleModels{
		DoctorID:    req.DoctorID,
		Date:        req.Date,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		IsAvailable: true,
		CreatedAt:   helpers.TimeStampNow(),
		UpdatedAt:   "",
	}

	t.Run("Failure Case - Error Create Schedule", func(t *testing.T) {
		expectedErr := errors.New("error creating schedules")

		ts.ScheduleRepo.On("CreateSchedule", newData).Return(int64(0), expectedErr).Once()

		result, err := ts.ScheduleService.CreateSchedule(req)

		assert.Equal(t, int64(0), result)
		assert.Error(t, err)

	})
}

func Test_FindByID(t *testing.T) {
	ts := test.SetupTestCase(t)

	req := models.RequestID{
		ID: 1,
	}

	t.Run("Success Case - Find By ID", func(t *testing.T) {
		expected := models.ScheduleModels{
			ID:          1,
			DoctorID:    1,
			Date:        "2022-01-01",
			StartTime:   "09:00",
			EndTime:     "12:00",
			IsAvailable: true,
			CreatedAt:   helpers.TimeStampNow(),
			UpdatedAt:   "",
		}

		ts.ScheduleRepo.On("FindScheduleByID", req.ID).Return(expected, nil).Once()

		result, err := ts.ScheduleService.FindByID(req)

		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("Failure Case - Find By ID", func(t *testing.T) {
		expectedErr := errors.New("schedule not founds")

		ts.ScheduleRepo.On("FindScheduleByID", req.ID).Return(models.ScheduleModels{}, expectedErr).Once()

		result, err := ts.ScheduleService.FindByID(req)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, models.ScheduleModels{}, result)
	})

}
