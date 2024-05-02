package sqldb

import (
	"context"
	"database/sql"
)

type SQLDBService interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type SQLDB struct {
	db *sql.DB
}

func NewSQLDB(db *sql.DB) *SQLDB {
	return &SQLDB{db}
}

func (s *SQLDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.QueryContext(ctx, query, args...)
}

func (s *SQLDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.db.ExecContext(ctx, query, args...)
}
