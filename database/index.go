package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("todo_app.db"))
	return db
}
