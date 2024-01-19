package service

import (
	"fmt"
	"time"

	"github.com/openkrafter/anytore-backend/model"
)

func getTrainingItemById(id int) model.TrainingItem {
	// DB Access
	trainingItem := new(model.TrainingItem)
	trainingItem.Id = id
	trainingItem.UserId = 1
	trainingItem.Name = "running"
	trainingItem.Type = "aerobic"
	trainingItem.Unit = "minute"
	trainingItem.Kcal = 2
	trainingItem.CreatedAt = time.Now()
	trainingItem.UpdatedAt = time.Now()

	return *trainingItem
}

func GetTraningItem(id int) model.TrainingItem {
	fmt.Println("trainingItem id: ", id)
	trainingItem := getTrainingItemById(id)
	return trainingItem
}
