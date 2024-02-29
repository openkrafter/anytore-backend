package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	anytoreConfig "github.com/openkrafter/anytore-backend/config"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
	testenvironment "github.com/openkrafter/anytore-backend/test-environment"
)

func setup() error {
	logger.InitLogger()
	logger.Logger.Info("controller package test setup")

	anytoreConfig.InitConfig()
	testenvironment.SetupDynamoDbClient()

	go Run()
    time.Sleep(10 * time.Millisecond)

	return nil
}

func teardown() error {
	logger.Logger.Info("controller package test done")
	return nil
}

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		logger.Logger.Error("Failed to setup.", logger.ErrAttr(err))
	}

	ret := m.Run()

	if err := teardown(); err != nil {
		logger.Logger.Error("Failed to teardown.", logger.ErrAttr(err))
	}

	os.Exit(ret)
}

func TestGetTraningItem(t *testing.T) {
	tests := []struct {
		name string
		dynamoInput *model.TrainingItem
		url string
		wantErr bool
	}{
		{
			name: "case1",
			dynamoInput: &model.TrainingItem{
				Id:     1,
				UserId: 1,
				Name:   "ランニング",
				Type:   "aerobic",
				Unit:   "hour",
				Kcal:   150,
			},
			url: "http://localhost:8080/training-items/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testenvironment.SetupTraningItemTestData(tt.dynamoInput)
			defer testenvironment.TeardownTraningItemTestData(tt.dynamoInput.Id)

			resp, err := http.Get(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTraningItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTraningItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var got map[string]interface{}
			json.Unmarshal(body, &got)

			var want map[string]interface{}
			wantBytes, err := json.Marshal(tt.dynamoInput)
			if err != nil {
				logger.Logger.Error("Failed to marshal.", logger.ErrAttr(err))
			}
			json.Unmarshal(wantBytes, &want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetTraningItem() = %v, want %v", got, want)
			}
		})
	}
}
