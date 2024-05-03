package mock

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/openkrafter/anytore-backend/database/sqldb"
)

type MockSQLDBService struct {
	Db   *sql.DB
	Mock sqlmock.Sqlmock
}

func NewMockSQLDBService() (*MockSQLDBService, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	sqldb.SQLDBClient = db
	return &MockSQLDBService{Db: db, Mock: mock}, nil
}

func CloseMockSQLDBService(s *MockSQLDBService) error {
	return s.Db.Close()
}

func (s *MockSQLDBService) QueryContext(query string, args ...interface{}) (*sql.Rows, error) {
	return s.Db.Query(query, args...)
}

func (s *MockSQLDBService) QueryRowContext(query string, args ...interface{}) *sql.Row {
	return s.Db.QueryRow(query, args...)
}

func (s *MockSQLDBService) ExecContext(query string, args ...interface{}) (sql.Result, error) {
	return s.Db.Exec(query, args...)
}
