package model

type TrainingRecord struct {
	Id             int     `json:"id"`
	UserId         int     `json:"userId"`
	TrainingItemId int     `json:"trainingItemId"`
	Record         float64 `json:"record"`
	Date           int64   `json:"date"`
}
