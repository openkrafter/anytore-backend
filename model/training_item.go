package model

import "time"

type TrainingItem struct {
	Id int
	UserId int
	Name string
	Type string
	Unit string
	Kcal int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ti *TrainingItem) GetResponse() map[string] interface {} {
	res := make(map[string]interface{})
	res["id"] = ti.Id
	res["user_id"] = ti.UserId
	res["name"] = ti.Name
	res["type"] = ti.Type
	res["unit"] = ti.Unit
	res["kcal"] = ti.Kcal

	return res
}
