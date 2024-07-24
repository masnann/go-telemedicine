package models

type ScheduleModels struct {
	ID          int64  `json:"id"`
	DoctorID    int64  `json:"doctorID"`
	Date        string `json:"date"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	IsAvailable bool   `json:"isAvailable"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeleteAt    string `json:"-"`
}

type ScheduleCreateRequest struct {
	DoctorID    int64  `json:"doctorID"`
	Date        string `json:"date"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	IsAvailable bool   `json:"isAvailable"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeleteAt    string `json:"-"`
}

type ScheduleFindListAvailableRequest struct {
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
