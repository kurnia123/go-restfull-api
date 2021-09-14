package app

import (
	"database/sql"
	"fmt"
	"kurnia123/restfull-api-v1/helper"
	"os"
	"time"

	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *sql.DB {
	connect_db := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := sql.Open("mysql", connect_db)
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func NewGorm(db *sql.DB) *gorm.DB {
	gorm, err := gorm.Open(gorm_mysql.New(gorm_mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	helper.PanicError(err)

	return gorm
}
