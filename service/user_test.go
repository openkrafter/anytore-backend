package service_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/openkrafter/anytore-backend/model"
	"github.com/openkrafter/anytore-backend/service"
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
					Id:        1,
					Name:      "user1",
					Email:     "user1@example.com",
					Password:  "password1",
					CreatedAt: time.Unix(1714648189, 0),
					UpdatedAt: time.Unix(1714648189, 0),
				},
				{
					Id:        2,
					Name:      "user2",
					Email:     "user2@example.com",
					Password:  "password2",
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
			defer func() {
				mockService.Mock.ExpectClose()
				if err := mock.CloseMockSQLDBService(mockService); err != nil {
					t.Fatalf("Failed to close mock service: %v", err)
				}
			}()

			query := `
				SELECT
					id,
					name,
					email,
					password,
					created_at,
					updated_at
				FROM
					users
			`

			rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
				AddRow(tt.want[0].Id, tt.want[0].Name, tt.want[0].Email, tt.want[0].Password, tt.want[0].CreatedAt, tt.want[0].UpdatedAt).
				AddRow(tt.want[1].Id, tt.want[1].Name, tt.want[1].Email, tt.want[1].Password, tt.want[1].CreatedAt, tt.want[1].UpdatedAt)

			mockService.Mock.ExpectQuery(query).WillReturnRows(rows)

			got, err := service.GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: &model.User{
				Id:        1,
				Name:      "user1",
				Email:     "user1@example.com",
				Password:  "password1",
				CreatedAt: time.Unix(1714648189, 0),
				UpdatedAt: time.Unix(1714648189, 0),
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
			defer func() {
				mockService.Mock.ExpectClose()
				if err := mock.CloseMockSQLDBService(mockService); err != nil {
					t.Fatalf("Failed to close mock service: %v", err)
				}
			}()

			query := `
				SELECT
					id,
					name,
					email,
					password,
					created_at,
					updated_at
				FROM
					users
				WHERE
					id = \?
			`

			row := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
				AddRow(tt.want.Id, tt.want.Name, tt.want.Email, tt.want.Password, tt.want.CreatedAt, tt.want.UpdatedAt)

			mockService.Mock.ExpectQuery(query).WithArgs(tt.args.id).WillReturnRows(row)

			got, err := service.GetUserById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Name:     "user1",
					Email:    "user1@example.com",
					Password: "password1",
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
			defer func() {
				mockService.Mock.ExpectClose()
				if err := mock.CloseMockSQLDBService(mockService); err != nil {
					t.Fatalf("Failed to close mock service: %v", err)
				}
			}()

			query := `
				INSERT INTO users \(
					name,
					email,
					password
				\) VALUES \(
					\?,
					\?,
					\?
				\)
			`

			mockHasher := mock.NewMockHasher()
			hashedPassword, _ := mockHasher.HashPassword(tt.args.user.Password)

			mockService.Mock.ExpectExec(query).
				WithArgs(tt.args.user.Name, tt.args.user.Email, hashedPassword).
				WillReturnResult(sqlmock.NewResult(1, 1))

			if err := service.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Id:       1,
					Name:     "user1",
					Email:    "user1@example.com",
					Password: "password1",
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
			defer func() {
				mockService.Mock.ExpectClose()
				if err := mock.CloseMockSQLDBService(mockService); err != nil {
					t.Fatalf("Failed to close mock service: %v", err)
				}
			}()

			query := `
				UPDATE users
				SET
					name = \?,
					email = \?,
					password = \?
				WHERE
					id = \?
			`

			mockHasher := mock.NewMockHasher()
			hashedPassword, _ := mockHasher.HashPassword(tt.args.user.Password)

			mockService.Mock.ExpectExec(query).
				WithArgs(tt.args.user.Name, tt.args.user.Email, hashedPassword, tt.args.user.Id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			if err := service.UpdateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				ctx: context.Background(),
				id:  1,
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
			defer func() {
				mockService.Mock.ExpectClose()
				if err := mock.CloseMockSQLDBService(mockService); err != nil {
					t.Fatalf("Failed to close mock service: %v", err)
				}
			}()

			query := `
				DELETE FROM users
				WHERE
					id = \?
			`

			mockService.Mock.ExpectExec(query).
				WithArgs(tt.args.id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			if err := service.DeleteUser(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
