package mocks

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// NewMongoDB returns a new MongoDB mock
func NewMongoDB(t *testing.T) *mtest.T {
	return mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
}

// NewSqlDB returns a new SQL DB mock
func NewSqlDB(t *testing.T) (sqlMock sqlmock.Sqlmock, db *sql.DB) {
	db, sqlMock, _ = sqlmock.New()
	return
}
