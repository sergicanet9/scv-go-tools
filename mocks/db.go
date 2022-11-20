package mocks

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func NewMongoDB(t *testing.T) *mtest.T {
	return mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
}

func NewSqlDB(t *testing.T) (sqlMock sqlmock.Sqlmock, db *sql.DB) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error creating the mock: %s", err)
	}
	return
}
