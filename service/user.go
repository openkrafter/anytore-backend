package service

import (
	"context"
	"strings"

	"github.com/openkrafter/anytore-backend/auth"
	"github.com/openkrafter/anytore-backend/database/sqldb"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
)

func GetUsers(ctx context.Context) ([]*model.User, error) {
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

	rows, err := sqldb.SQLDBClient.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func GetUser(ctx context.Context, id int) (*model.User, error) {
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
			id = ?
	`

	row := sqldb.SQLDBClient.QueryRowContext(ctx, query, id)

	var user model.User
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (
			name,
			email,
			password
		) VALUES (
			?,
			?,
			?
		)
	`

	hashedPassword, err := auth.PassHasher.HashPassword(user.Password)
	if err != nil {
		logger.Logger.Error("Error hashing password", err)
		return err
	}

	_, err = sqldb.SQLDBClient.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		hashedPassword,
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(ctx context.Context, user *model.User) error {
	var params []interface{}
	setClause := []string{}

	if user.Name != "" {
		setClause = append(setClause, "name = ?")
		params = append(params, user.Name)
	}

	if user.Email != "" {
		setClause = append(setClause, "email = ?")
		params = append(params, user.Email)
	}

	if user.Password != "" {
		hashedPassword, err := auth.PassHasher.HashPassword(user.Password)
		if err != nil {
			logger.Logger.Error("Error hashing password", err)
			return err
		}
		setClause = append(setClause, "password = ?")
		params = append(params, hashedPassword)
	}

	if len(setClause) == 0 {
		return nil
	}

	params = append(params, user.Id)

	query := `
		UPDATE users
		SET ` + strings.Join(setClause, ", ") + `
		WHERE
			id = ?
	`

	_, err := sqldb.SQLDBClient.ExecContext(
		ctx,
		query,
		params...,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(ctx context.Context, id int) error {
	query := `
		DELETE FROM users
		WHERE
			id = ?
	`

	_, err := sqldb.SQLDBClient.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
