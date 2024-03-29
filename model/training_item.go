package model

type TrainingItem struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Unit string `json:"unit"`
	Kcal int `json:"kcal"`
}

func (ti *TrainingItem) GetResponse() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = ti.Id
	res["user_id"] = ti.UserId
	res["name"] = ti.Name
	res["type"] = ti.Type
	res["unit"] = ti.Unit
	res["kcal"] = ti.Kcal

	return res
}
