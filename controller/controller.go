package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
	"github.com/openkrafter/anytore-backend/service"
)

func SampleTraningItem(c *gin.Context) {
	trainingItem := new(model.TrainingItem)
	trainingItem.Id = 10
	trainingItem.UserId = 1
	trainingItem.Name = "running"
	trainingItem.Type = "aerobic"
	trainingItem.Unit = "hour"
	trainingItem.Kcal = 100

	response := trainingItem.GetResponse()
	c.JSON(http.StatusOK, response)
}

func GetTraningItem(c *gin.Context) {
	trainingItemId, _ := strconv.Atoi(c.Param("training-item-id"))
	trainingItem, err := service.GetTraningItem(trainingItemId)
	if err != nil {
		logger.Logger.Error("GetTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	response := trainingItem.GetResponse()
	c.JSON(http.StatusOK, response)
}

func Run() {
	logger.Logger.Info("Controller thread start.")

	r := gin.Default()
	r.GET("/sample", SampleTraningItem)
	r.GET("/training-items/:training-item-id", GetTraningItem)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
