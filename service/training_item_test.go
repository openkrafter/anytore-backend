package service

import (
	"log"
	"reflect"
	"sort"
	"testing"

	"github.com/openkrafter/anytore-backend/model"
	testenvironment "github.com/openkrafter/anytore-backend/test/environment"
)

// TODO: Add test cases.

func TestGetTraningItems(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name            string
		setupDynamoData []*model.TrainingItem
		args            args
		want            []*model.TrainingItem
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
			args: args{userId: 1},
			want: []*model.TrainingItem{
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

			got, err := GetTraningItems(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTraningItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sort.Slice(got, func(i, j int) bool {
				compA := *got[i]
				compB := *got[j]
				return compA.Id < compB.Id
			})

			if !reflect.DeepEqual(got, tt.want) {
				for _, trainingItem := range got {
					log.Println(trainingItem)
				}
				for _, trainingItem := range tt.want {
					log.Println(trainingItem)
				}

				t.Errorf("GetTraningItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTraningItem(t *testing.T) {
	type args struct {
		id     int
		userId int
	}
	tests := []struct {
		name            string
		setupDynamoData []*model.TrainingItem
		args            args
		want            *model.TrainingItem
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
			args: args{id: 1, userId: 1},
			want: &model.TrainingItem{
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

			got, err := GetTraningItem(tt.args.id, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTraningItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTraningItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIncrementId(t *testing.T) {
	tests := []struct {
		name    string
		want    reflect.Type
		wantErr bool
	}{
		{
			name: "case1",
			want: reflect.TypeOf(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			incrementId, err := GetIncrementId()
			got := reflect.TypeOf(incrementId)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetIncrementId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIncrementId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateTraningItem(t *testing.T) {
	type args struct {
		input  *model.TrainingItem
		userId int
	}
	tests := []struct {
		name            string
		setupDynamoData []*model.TrainingItem
		args            args
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
					Kcal:   120,
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
				input: &model.TrainingItem{
					Id:     1,
					UserId: 1,
					Name:   "ランニング",
					Type:   "aerobic",
					Unit:   "hour",
					Kcal:   150,
				},
				userId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, trainingItem := range tt.setupDynamoData {
				testenvironment.SetupTraningItemTestData(trainingItem)
				defer testenvironment.TeardownTraningItemTestData(trainingItem.Id)
			}

			if err := UpdateTraningItem(tt.args.input, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTraningItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTraningItem(t *testing.T) {
	type args struct {
		id     int
		userId int
	}
	tests := []struct {
		name            string
		setupDynamoData []*model.TrainingItem
		args            args
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
					Kcal:   120,
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
			args: args{id: 1, userId: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, trainingItem := range tt.setupDynamoData {
				testenvironment.SetupTraningItemTestData(trainingItem)
				defer testenvironment.TeardownTraningItemTestData(trainingItem.Id)
			}

			if err := DeleteTraningItem(tt.args.id, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTraningItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
