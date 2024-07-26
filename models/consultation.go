package models

type ConsultationModels struct {
	ID               int64  `json:"id"`
	PatientID        int64  `json:"patientID"`
	PatientName      string `json:"patientName"`
	PatientEmail     string `json:"patientEmail"`
	ScheduleID       int64  `json:"scheduleID"`
	DoctorID         int64  `json:"doctorID"`
	DoctorName       string `json:"doctorName"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	ConsultationType string `json:"consultationType"`
	Status           string `json:"status"`
	Notes            string `json:"notes"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
	DeletedAt        string `json:"-"`
}

type ConsultationCreateRequest struct {
	PatientID        int64  `json:"patientID"`
	PatientName      string `json:"patientName"`
	PatientEmail     string `json:"patientEmail"`
	ScheduleID       int64  `json:"scheduleID"`
	DoctorID         int64  `json:"doctorID"`
	DoctorName       string `json:"doctorName"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	ConsultationType string `json:"consultationType"`
	Status           string `json:"status"`
	Notes            string `json:"notes"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
	DeletedAt        string `json:"-"`
}

type ConsultationFindListByPatientIDRequest struct {
	PatientID  int64      `json:"patientID"`
	DateStart  string     `json:"dateStart"`
	DateEnd    string     `json:"dateEnd"`
	Status     string     `json:"status"`
	Pagination Pagination `json:"pagination"`
}

type ConsultationFindListByDoctorIDRequest struct {
	DoctorID  int64      `json:"doctorID"`
	DateStart  string     `json:"dateStart"`
	DateEnd    string     `json:"dateEnd"`
	Status     string     `json:"status"`
	Pagination Pagination `json:"pagination"`
}
