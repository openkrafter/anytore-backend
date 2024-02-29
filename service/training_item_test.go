package service

import (
	"reflect"
	"testing"

	"github.com/openkrafter/anytore-backend/model"
	testenvironment "github.com/openkrafter/anytore-backend/test-environment"
)

func TestGetTraningItem(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.TrainingItem
		wantErr bool
	}{
		{
			name: "case1",
			args: args{id: 1},
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
			testenvironment.SetupTraningItemTestData(tt.want)
			defer testenvironment.TeardownTraningItemTestData(tt.want.Id)

			got, err := GetTraningItem(tt.args.id)
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
