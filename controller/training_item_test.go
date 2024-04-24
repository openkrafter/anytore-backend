package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/model"
	testenvironment "github.com/openkrafter/anytore-backend/test/environment"
)

func TestGetTraningItems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type args struct {
		path   string
		userId int
	}
	tests := []struct {
		name            string
		setupDynamoData []*model.TrainingItem
		args            args
		wantStatusCode  int
		wantBody        []model.TrainingItem
		wantErr         bool
	}{
		{
			name: "case1",
			setupDynamoData: []*model.TrainingItem{
				{
					Id:     1,
					UserId: 1,
					Name:   "ランニング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   150,
				},
				{
					Id:     3,
					UserId: 1,
					Name:   "ウォーキング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   90,
				},
				{
					Id:     2,
					UserId: 2,
					Name:   "ウォーキング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   100,
				},
			},
			args: args{
				path:   "/training-items",
				userId: 1,
			},
			wantStatusCode: http.StatusOK,
			wantBody: []model.TrainingItem{
				{
					Id:     1,
					UserId: 1,
					Name:   "ランニング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   150,
				},
				{
					Id:     3,
					UserId: 1,
					Name:   "ウォーキング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   90,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, trainingItem := range tt.setupDynamoData {
				testenvironment.SetupTraningItemTestData(trainingItem)
				defer testenvironment.TeardownTraningItemTestData(trainingItem.Id)
			}

			r := gin.Default()
			RegisterRoutes(r)
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", tt.args.path, nil)
			if err != nil {
				t.Errorf("NewRequest error = %v", err)
				return
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %d", tt.args.userId))

			r.ServeHTTP(w, req)
			if w.Code != tt.wantStatusCode {
				t.Errorf("GetTraningItems API Status code = %d, want %d", w.Code, tt.wantStatusCode)
			}

			var got []model.TrainingItem
			err = json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			if !reflect.DeepEqual(got, tt.wantBody) {
				t.Errorf("GetTraningItems API = %v, want %v", got, tt.wantBody)
			}
		})
	}
}

func TestGetTraningItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type args struct {
		path   string
		userId int
	}
	tests := []struct {
		name            string
		setupDynamoData []*model.TrainingItem
		args            args
		wantStatusCode  int
		wantBody        model.TrainingItem
		wantErr         bool
	}{
		{
			name: "case1",
			setupDynamoData: []*model.TrainingItem{
				{
					Id:     1,
					UserId: 1,
					Name:   "ランニング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   150,
				},
				{
					Id:     2,
					UserId: 2,
					Name:   "ウォーキング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   100,
				},
			},
			args: args{
				path:   "/training-items/1",
				userId: 1,
			},
			wantStatusCode: http.StatusOK,
			wantBody: model.TrainingItem{
				Id:     1,
				UserId: 1,
				Name:   "ランニング",
				Type:   "aerobic",
				Unit:   "hour",
				Kcal:   150,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, trainingItem := range tt.setupDynamoData {
				testenvironment.SetupTraningItemTestData(trainingItem)
				defer testenvironment.TeardownTraningItemTestData(trainingItem.Id)
			}

			r := gin.Default()
			RegisterRoutes(r)
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", tt.args.path, nil)
			if err != nil {
				t.Errorf("NewRequest error = %v", err)
				return
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %d", tt.args.userId))

			r.ServeHTTP(w, req)
			if w.Code != tt.wantStatusCode {
				t.Errorf("GetTraningItem API Status code = %d, want %d", w.Code, tt.wantStatusCode)
			}

			var got model.TrainingItem
			err = json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			// log.Println(got)
			if !reflect.DeepEqual(got, tt.wantBody) {
				t.Errorf("GetTraningItem API = %v, want %v", got, tt.wantBody)
			}
		})
	}
}

// func TestCreateTraningItem(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		setupDynamoData *model.TrainingItem
// 		url         string
// 		want		map[string]interface{}
// 		wantErr     bool
// 	}{
// 		{
// 			name: "case1",
// 			setupDynamoData: &model.TrainingItem{
// 				UserId: 1,
// 				Name:   "ランニング",
// 				Type:   "aerobic",
// 				Unit:   "hour",
// 				Kcal:   150,
// 			},
// 			url: "http://localhost:8080/training-items",
// 			want: map[string]interface{}{
// 				"status": http.StatusCreated,
// 				"body": &model.TrainingItem{
// 					UserId: 1,
// 					Name:   "ランニング",
// 					Type:   "aerobic",
// 					Unit:   "hour",
// 					Kcal:   150,
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			client := &http.Client{
// 				CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 					return http.ErrUseLastResponse
// 				},
// 			}
//
// 			requestBody, err := json.Marshal(tt.setupDynamoData)
// 			if err != nil {
// 				logger.Logger.Error("Failed to marshal.", logger.ErrAttr(err))
// 				return
// 			}
//
// 			req, err := http.NewRequest(
// 				"POST",
// 				tt.url,
// 				bytes.NewBuffer(requestBody))
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreateTraningItem API error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
//
// 			res, err := client.Do(req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreateTraningItem API error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			defer res.Body.Close()
//
// 			body, err := io.ReadAll(res.Body)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetTraningItem API error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			got := map[string]interface{}{
// 				"status": res.StatusCode,
// 				"body": &model.TrainingItem{},
// 			}
// 			json.Unmarshal(body, got["body"])
//
// 			gotBody := got["body"].(*model.TrainingItem)
// 			wantBody := tt.want["body"].(*model.TrainingItem)
// 			wantBody.Id = gotBody.Id
// 			defer testenvironment.TeardownTraningItemTestData(gotBody.Id)
//
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("CreateTraningItem API = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestUpdateTraningItem(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		setupDynamoData *model.TrainingItem
// 		url         string
// 		want		map[string]interface{}
// 		wantErr     bool
// 	}{
// 		{
// 			name: "case1",
// 			setupDynamoData: &model.TrainingItem{
// 				Id:     1,
// 				UserId: 1,
// 				Name:   "ランニング",
// 				Type:   "aerobic",
// 				Unit:   "hour",
// 				Kcal:   150,
// 			},
// 			url: "http://localhost:8080/training-items/1",
// 			want: map[string]interface{}{
// 				"status": http.StatusCreated,
// 				"body": &model.TrainingItem{
// 					Id:     1,
// 					UserId: 1,
// 					Name:   "ランニング",
// 					Type:   "aerobic",
// 					Unit:   "hour",
// 					Kcal:   150,
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			defer testenvironment.TeardownTraningItemTestData(tt.setupDynamoData.Id)
//
// 			client := &http.Client{
// 				CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 					return http.ErrUseLastResponse
// 				},
// 			}
//
// 			requestBody, err := json.Marshal(tt.setupDynamoData)
// 			if err != nil {
// 				logger.Logger.Error("Failed to marshal.", logger.ErrAttr(err))
// 				return
// 			}
//
// 			req, err := http.NewRequest(
// 				"PUT",
// 				tt.url,
// 				bytes.NewBuffer(requestBody))
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UpdateTraningItem API error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
//
// 			res, err := client.Do(req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UpdateTraningItem API error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			defer res.Body.Close()
//
// 			body, err := io.ReadAll(res.Body)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetTraningItem API error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			got := map[string]interface{}{
// 				"status": res.StatusCode,
// 				"body": &model.TrainingItem{},
// 			}
// 			json.Unmarshal(body, got["body"])
//
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("DeleteTraningItem API = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
// func TestDeleteTraningItem(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		setupDynamoData *model.TrainingItem
// 		url         string
// 		want		map[string]interface{}
// 		wantErr     bool
// 	}{
// 		{
// 			name: "case1",
// 			setupDynamoData: &model.TrainingItem{
// 				Id:     1,
// 				UserId: 1,
// 				Name:   "ランニング",
// 				Type:   "aerobic",
// 				Unit:   "hour",
// 				Kcal:   150,
// 			},
// 			url: "http://localhost:8080/training-items/1",
// 			want: map[string]interface{}{
// 				"status": http.StatusNoContent,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			testenvironment.SetupTraningItemTestData(tt.setupDynamoData)
// 			defer testenvironment.TeardownTraningItemTestData(tt.setupDynamoData.Id)
//
// 			client := &http.Client{
// 				CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 					return http.ErrUseLastResponse
// 				},
// 			}
//
// 			req, err := http.NewRequest("DELETE", tt.url, nil)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DeleteTraningItem API error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
//
// 			res, err := client.Do(req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DeleteTraningItem API error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			defer res.Body.Close()
//
// 			got := map[string]interface{}{
// 				"status": res.StatusCode,
// 			}
//
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("DeleteTraningItem API = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
