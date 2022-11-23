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
	db, sqlMock, _ = sqlmock.New()
	return
}
