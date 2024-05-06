package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/controller"
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
				if err := testenvironment.SetupTraningItemTestData(trainingItem); err != nil {
					t.Fatalf("SetupTraningItemTestData Faled: %v", err)
				}
			}
			defer func() {
				if err := testenvironment.TeardownTraningItemTestData(); err != nil {
					t.Fatalf("TeardownTraningItemTestData Faled: %v", err)
				}
			}()

			r := gin.Default()
			controller.RegisterRoutes(r)
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", tt.args.path, nil)
			if err != nil {
				t.Fatalf("NewRequest error = %v", err)
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
				if err := testenvironment.SetupTraningItemTestData(trainingItem); err != nil {
					t.Fatalf("SetupTraningItemTestData Faled: %v", err)
				}
			}
			defer func() {
				if err := testenvironment.TeardownTraningItemTestData(); err != nil {
					t.Fatalf("TeardownTraningItemTestData Faled: %v", err)
				}
			}()

			r := gin.Default()
			controller.RegisterRoutes(r)
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", tt.args.path, nil)
			if err != nil {
				t.Fatalf("NewRequest error = %v", err)
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

			if !reflect.DeepEqual(got, tt.wantBody) {
				t.Errorf("GetTraningItem API = %v, want %v", got, tt.wantBody)
			}
		})
	}
}

func TestCreateTraningItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type args struct {
		path         string
		userId       int
		trainingItem model.TrainingItem
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantBody       model.TrainingItem
		wantErr        bool
	}{
		{
			name: "case1",
			args: args{
				path:   "/training-items",
				userId: 1,
				trainingItem: model.TrainingItem{
					UserId: 1,
					Name:   "ランニング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   150,
				},
			},
			wantStatusCode: http.StatusCreated,
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
			defer func() {
				if err := testenvironment.TeardownTraningItemTestData(); err != nil {
					t.Fatalf("TeardownTraningItemTestData Faled: %v", err)
				}
			}()

			r := gin.Default()
			controller.RegisterRoutes(r)
			w := httptest.NewRecorder()

			jsonData, err := json.Marshal(tt.args.trainingItem)
			if err != nil {
				t.Errorf("json Marshal Error: %v", err)
				return
			}

			req, err := http.NewRequest("POST", tt.args.path, bytes.NewBuffer(jsonData))
			if err != nil {
				t.Errorf("NewRequest error = %v", err)
				return
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %d", tt.args.userId))
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)
			if w.Code != tt.wantStatusCode {
				t.Errorf("CreateTraningItem API Status code = %d, want %d", w.Code, tt.wantStatusCode)
			}

			var got model.TrainingItem
			err = json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			log.Println(got)
			if !reflect.DeepEqual(got, tt.wantBody) {
				t.Errorf("CreateTraningItem API = %v, want %v", got, tt.wantBody)
			}
		})
	}

}

func TestUpdateTraningItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type args struct {
		path         string
		userId       int
		trainingItem model.TrainingItem
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
				trainingItem: model.TrainingItem{
					Id:     1,
					UserId: 1,
					Name:   "ウォーキング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   90,
				},
			},
			wantStatusCode: http.StatusCreated,
			wantBody: model.TrainingItem{
				Id:     1,
				UserId: 1,
				Name:   "ウォーキング",
				Type:   "aerobic",
				Unit:   "hour",
				Kcal:   90,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, trainingItem := range tt.setupDynamoData {
				if err := testenvironment.SetupTraningItemTestData(trainingItem); err != nil {
					t.Fatalf("SetupTraningItemTestData Faled: %v", err)
				}
			}
			defer func() {
				if err := testenvironment.TeardownTraningItemTestData(); err != nil {
					t.Fatalf("TeardownTraningItemTestData Faled: %v", err)
				}
			}()

			r := gin.Default()
			controller.RegisterRoutes(r)
			w := httptest.NewRecorder()

			jsonData, err := json.Marshal(tt.args.trainingItem)
			if err != nil {
				t.Fatalf("json Marshal Error: %v", err)
			}

			req, err := http.NewRequest("PUT", tt.args.path, bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("NewRequest error = %v", err)
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %d", tt.args.userId))
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)
			if w.Code != tt.wantStatusCode {
				t.Errorf("UpdateTraningItem API Status code = %d, want %d", w.Code, tt.wantStatusCode)
			}

			var got model.TrainingItem
			err = json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			if !reflect.DeepEqual(got, tt.wantBody) {
				t.Errorf("UpdateTraningItem API = %v, want %v", got, tt.wantBody)
			}
		})
	}
}

func TestDeleteTraningItem(t *testing.T) {
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
			wantStatusCode: http.StatusNoContent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, trainingItem := range tt.setupDynamoData {
				if err := testenvironment.SetupTraningItemTestData(trainingItem); err != nil {
					t.Fatalf("SetupTraningItemTestData Faled: %v", err)
				}
			}
			defer func() {
				if err := testenvironment.TeardownTraningItemTestData(); err != nil {
					t.Fatalf("TeardownTraningItemTestData Faled: %v", err)
				}
			}()

			r := gin.Default()
			controller.RegisterRoutes(r)
			w := httptest.NewRecorder()

			req, err := http.NewRequest("DELETE", tt.args.path, nil)
			if err != nil {
				t.Fatalf("NewRequest error = %v", err)
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %d", tt.args.userId))

			r.ServeHTTP(w, req)
			if w.Code != tt.wantStatusCode {
				t.Errorf("DeleteTraningItem API Status code = %d, want %d", w.Code, tt.wantStatusCode)
			}

			if w.Body.String() != "" {
				t.Errorf("Unexpected body = %v, want an empty body", w.Body.String())
			}
		})
	}

}
