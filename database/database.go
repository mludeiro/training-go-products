package database

import (
	"log"
	"training-go-products/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	gormDB *gorm.DB
}

func (db *Database) InitializeSqlite() *Database {

	DB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db.gormDB = DB
	return db
}

func (db *Database) GetDB() *gorm.DB {
	return db.gormDB
}

func (db *Database) Migrate() *Database {
	db.GetDB().AutoMigrate(&entity.Product{})
	return db
}

func (db *Database) CreateTestData() *Database {
	productTest := entity.Product{Name: "HelloImYellow"}

	db.GetDB().Create(&productTest)

	return db
}
