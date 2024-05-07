package model

type TrainingItem struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Unit   string `json:"unit"`
	Kcal   int    `json:"kcal"`
}
