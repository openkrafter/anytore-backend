package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/config"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/service"
)

func GetTraningItem(c *gin.Context) {
	trainingItemId, _ := strconv.Atoi(c.Param("training-item-id"))
	trainingItem := service.GetTraningItem(trainingItemId)
	response := trainingItem.GetResponse()
	c.JSON(http.StatusOK, response)
}

func Run() {
	logger.Logger.Info("Controller thread start.")

	logger.Logger.Info(config.Config.GIN_MODE)
	r := gin.Default()
	r.GET("/training-items/:training-item-id", GetTraningItem)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
