package model

import (
	"backend/util"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB database connection
var DB *gorm.DB

// Database init postgres connection
func Database(connString string) {
	db, err := gorm.Open("postgres", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("database connection failed", err)
	}
	//setting connection pool
	//idle
	db.DB().SetMaxIdleConns(50)
	//ope
	db.DB().SetMaxOpenConns(100)
	//timeout
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
