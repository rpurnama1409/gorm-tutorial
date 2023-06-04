package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetGormConn() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=beranda3525 dbname=gorm-tutorial port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err

}
