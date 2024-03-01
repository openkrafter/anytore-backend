package controller

import (
	"bytes"
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
		name        string
		dynamoInput *model.TrainingItem
		url         string
		want		map[string]interface{}
		wantErr     bool
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
			want: map[string]interface{}{
				"status": http.StatusOK,
				"body": &model.TrainingItem{
					Id:     1,
					UserId: 1,
					Name:   "ランニング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   150,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testenvironment.SetupTraningItemTestData(tt.dynamoInput)
			defer testenvironment.TeardownTraningItemTestData(tt.dynamoInput.Id)

			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			req, err := http.NewRequest("GET", tt.url, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			res, err := client.Do(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := map[string]interface{}{
				"status": res.StatusCode,
				"body": &model.TrainingItem{},
			}
			json.Unmarshal(body, got["body"])

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTraningItem API = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTraningItem(t *testing.T) {
	tests := []struct {
		name        string
		dynamoInput *model.TrainingItem
		url         string
		want		map[string]interface{}
		wantErr     bool
	}{
		{
			name: "case1",
			dynamoInput: &model.TrainingItem{
				UserId: 1,
				Name:   "ランニング",
				Type:   "aerobic",
				Unit:   "hour",
				Kcal:   150,
			},
			url: "http://localhost:8080/training-items",
			want: map[string]interface{}{
				"status": http.StatusCreated,
				"body": &model.TrainingItem{
					UserId: 1,
					Name:   "ランニング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   150,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			requestBody, err := json.Marshal(tt.dynamoInput)
			if err != nil {
				logger.Logger.Error("Failed to marshal.", logger.ErrAttr(err))
				return
			}

			req, err := http.NewRequest(
				"POST",
				tt.url,
				bytes.NewBuffer(requestBody))
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			res, err := client.Do(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := map[string]interface{}{
				"status": res.StatusCode,
				"body": &model.TrainingItem{},
			}
			json.Unmarshal(body, got["body"])

			gotBody := got["body"].(*model.TrainingItem)
			wantBody := tt.want["body"].(*model.TrainingItem)
			wantBody.Id = gotBody.Id
			defer testenvironment.TeardownTraningItemTestData(gotBody.Id)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTraningItem API = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateTraningItem(t *testing.T) {
	tests := []struct {
		name        string
		dynamoInput *model.TrainingItem
		url         string
		want		map[string]interface{}
		wantErr     bool
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
			want: map[string]interface{}{
				"status": http.StatusCreated,
				"body": &model.TrainingItem{
					Id:     1,
					UserId: 1,
					Name:   "ランニング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   150,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testenvironment.TeardownTraningItemTestData(tt.dynamoInput.Id)

			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			requestBody, err := json.Marshal(tt.dynamoInput)
			if err != nil {
				logger.Logger.Error("Failed to marshal.", logger.ErrAttr(err))
				return
			}

			req, err := http.NewRequest(
				"PUT",
				tt.url,
				bytes.NewBuffer(requestBody))
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			res, err := client.Do(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := map[string]interface{}{
				"status": res.StatusCode,
				"body": &model.TrainingItem{},
			}
			json.Unmarshal(body, got["body"])

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTraningItem API = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDeleteTraningItem(t *testing.T) {
	tests := []struct {
		name        string
		dynamoInput *model.TrainingItem
		url         string
		want		map[string]interface{}
		wantErr     bool
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
			want: map[string]interface{}{
				"status": http.StatusNoContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testenvironment.SetupTraningItemTestData(tt.dynamoInput)
			defer testenvironment.TeardownTraningItemTestData(tt.dynamoInput.Id)

			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			req, err := http.NewRequest("DELETE", tt.url, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			res, err := client.Do(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTraningItem API error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer res.Body.Close()

			got := map[string]interface{}{
				"status": res.StatusCode,
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTraningItem API = %v, want %v", got, tt.want)
			}
		})
	}
}