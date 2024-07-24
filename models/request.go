package models

type RequestID struct {
	ID int64 `json:"id"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
