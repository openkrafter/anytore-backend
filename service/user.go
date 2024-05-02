package service

import (
	"context"

	"github.com/openkrafter/anytore-backend/database/sqldb"
	"github.com/openkrafter/anytore-backend/model"
)

func GetUsers(ctx context.Context) ([]*model.User, error) {
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

	rows, err := sqldb.SQLDBClient.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Salt,
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

func CreateUser(ctx context.Context, user *model.User) error {
	// db, err := sql.Open("mysql", "develop:example@/anytore")
	// if err != nil {
	// 	return err
	// }
	// defer db.Close()

	query := `
		INSERT INTO users (
			name,
			email,
			password,
			salt
		) VALUES (
			?,
			?,
			?,
			?
		)
	`

	// _, err = db.ExecContext(
	_, err := sqldb.SQLDBClient.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Salt,
	)
	if err != nil {
		return err
	}

	return nil
}
