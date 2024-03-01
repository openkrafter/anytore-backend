package controller

import (
	"encoding/json"
	"io"
	"log"
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

func ListTraningItem(c *gin.Context) {
	trainingItems, err := service.GetTraningItems()
	if err != nil {
		logger.Logger.Error("ListTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	var response []map[string]interface{}
	for _, trainingItem := range trainingItems {
		response = append(response, trainingItem.GetResponse())
	}
	log.Println(response)
	c.JSON(http.StatusOK, response)
}

func GetTraningItem(c *gin.Context) {
	trainingItemId, err := strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("GetTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	trainingItem, err := service.GetTraningItem(trainingItemId)
	if err != nil {
		logger.Logger.Error("GetTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	response := trainingItem.GetResponse()
	c.JSON(http.StatusOK, response)
}

func CreateTraningItem(c *gin.Context) {
	var requestBody model.TrainingItem
	requestBodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	err = json.Unmarshal(requestBodyBytes, &requestBody)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	requestBody.Id, err = service.GetIncrementId()
	if err != nil {
		logger.Logger.Error("CreateTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	
	err = service.UpdateTraningItem(&requestBody)
	if err != nil {
		logger.Logger.Error("CreateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody.GetResponse())
}

func UpdateTraningItem(c *gin.Context) {
	var requestBody model.TrainingItem
	requestBodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	err = json.Unmarshal(requestBodyBytes, &requestBody)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	requestBody.Id, err = strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	
	err = service.UpdateTraningItem(&requestBody)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody.GetResponse())
}

func DeleteTraningItem(c *gin.Context) {
	trainingItemId, err := strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("DeleteTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	err = service.DeleteTraningItem(trainingItemId)
	if err != nil {
		logger.Logger.Error("DeleteTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	c.JSON(http.StatusNoContent, "Delete to update.")
}

func Run() {
	logger.Logger.Info("Controller thread start.")

	r := gin.Default()
	r.GET("/sample", SampleTraningItem) // for debug
	r.GET("/training-items", ListTraningItem)
	r.GET("/training-items/:training-item-id", GetTraningItem)
	r.POST("/training-items", CreateTraningItem)
	r.PUT("/training-items/:training-item-id", UpdateTraningItem)
	r.DELETE("/training-items/:training-item-id", DeleteTraningItem)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
