package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/customerror"
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
	userIdString := GetTokenFromAuthorizationHeader(c)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("Failed to convert userId int to int.", logger.ErrAttr(err))
		return
	}

	trainingItems, err := service.GetTraningItems(userId)
	if err != nil {
		logger.Logger.Error("ListTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	if trainingItems == nil {
		error404 := customerror.NewError404()
		c.JSON(error404.ErrorCode, error404.Body)
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
	userIdString := GetTokenFromAuthorizationHeader(c)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("Failed to convert userId int to int.", logger.ErrAttr(err))
		return
	}

	trainingItemId, err := strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("GetTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	trainingItem, err := service.GetTraningItem(trainingItemId, userId)
	if err != nil {
		logger.Logger.Error("GetTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	if trainingItem == nil {
		error404 := customerror.NewError404()
		c.JSON(error404.ErrorCode, error404.Body)
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
	
	err = service.CreateTraningItem(&requestBody)
	if err != nil {
		logger.Logger.Error("CreateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody.GetResponse())
}

func UpdateTraningItem(c *gin.Context) {
	userIdString := GetTokenFromAuthorizationHeader(c)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("Failed to convert userId int to int.", logger.ErrAttr(err))
		return
	}

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
	
	err = service.UpdateTraningItem(&requestBody, userId)
	if err != nil {
		if customErr, ok := err.(*customerror.Error404); ok {
			c.JSON(customErr.ErrorCode, customErr.Body)
			logger.Logger.Error("UpdateTraningItem 404.", logger.ErrAttr(err))
		} else {
			logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		}
		return
	}

	c.JSON(http.StatusCreated, requestBody.GetResponse())
}

func DeleteTraningItem(c *gin.Context) {
	userIdString := GetTokenFromAuthorizationHeader(c)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("Failed to convert userId int to int.", logger.ErrAttr(err))
		return
	}

	trainingItemId, err := strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("DeleteTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	err = service.DeleteTraningItem(trainingItemId, userId)
	if err != nil {
		if customErr, ok := err.(*customerror.Error404); ok {
			c.JSON(customErr.ErrorCode, customErr.Body)
			logger.Logger.Error("DeleteTraningItem 404.", logger.ErrAttr(err))
		} else {
			logger.Logger.Error("DeleteTraningItem Failed.", logger.ErrAttr(err))
		}
		return
	}
	c.JSON(http.StatusNoContent, "Delete to update.")
}

func Run() {
	logger.Logger.Info("Controller thread start.")

	r := gin.Default()

	if os.Getenv("GIN_MODE")  == "release" {
		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{"https://anytore.net"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Accept"},
			ExposeHeaders: []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge: 6 * time.Hour,
		}))
	} else {
		logger.Logger.Debug("CORS setting: debug mode")

		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:5173"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Accept"},
			ExposeHeaders: []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge: 6 * time.Hour,
		}))
	}

	r.GET("/sample", SampleTraningItem) // for debug
	r.GET("/training-items", ListTraningItem)
	r.GET("/training-items/:training-item-id", GetTraningItem)
	r.POST("/training-items", CreateTraningItem)
	r.PUT("/training-items/:training-item-id", UpdateTraningItem)
	r.DELETE("/training-items/:training-item-id", DeleteTraningItem)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
