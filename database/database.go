package database

import (
	"log"
	"time"
	"training-go-products/entity"
	"training-go-products/tools"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	gormDB *gorm.DB
}

func (db *Database) InitializeMySql() *Database {
	dsn := "root:pleasedonotstealmypassword(127.0.0.1:3306)/insertdbnamehere?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), createGormConfig())

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db.gormDB = DB
	return db
}

func (db *Database) Migrate() *Database {
	db.gormDB.AutoMigrate(&entity.Product{})
	return db
}

func (db *Database) GetDB() *gorm.DB {
	return db.gormDB
}

func (db *Database) CreateTestData() *Database {
	productTest := entity.Product{Name: "HelloImYellow"}

	db.gormDB.Create(&productTest)

	return db
}

func createGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(
			tools.GetLogger(), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			}),
	}
}
