package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//ConnectPsqlDB connect to PostgresDB
func ConnectPsqlDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
