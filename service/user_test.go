package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/openkrafter/anytore-backend/model"
	"github.com/openkrafter/anytore-backend/test/mock"
)

func TestGetUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.User
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				ctx: context.Background(),
			},
			want: []*model.User{
				{
					ID:        1,
					Name:      "user1",
					Email:     "user1@example.com",
					Password:  "password1",
					Salt:      "salt1",
					CreatedAt: time.Unix(1714648189, 0),
					UpdatedAt: time.Unix(1714648189, 0),
				},
				{
					ID:        2,
					Name:      "user2",
					Email:     "user2@example.com",
					Password:  "password2",
					Salt:      "salt2",
					CreatedAt: time.Unix(1714648289, 0),
					UpdatedAt: time.Unix(1714648289, 0),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService, err := mock.NewMockSQLDBService()
			if err != nil {
				t.Fatalf("Failed to create mock service: %v", err)
			}
			defer mock.CloseMockSQLDBService(mockService)

			query := `
				SELECT
					id,
					name,
					email,
					password,
					salt,
					created_at,
					updated_at
				FROM
					users
			`

			rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "salt", "created_at", "updated_at"}).
				AddRow(tt.want[0].ID, tt.want[0].Name, tt.want[0].Email, tt.want[0].Password, tt.want[0].Salt, tt.want[0].CreatedAt, tt.want[0].UpdatedAt).
				AddRow(tt.want[1].ID, tt.want[1].Name, tt.want[1].Email, tt.want[1].Password, tt.want[1].Salt, tt.want[1].CreatedAt, tt.want[1].UpdatedAt)

			mockService.Mock.ExpectQuery(query).WillReturnRows(rows)

			got, err := GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
