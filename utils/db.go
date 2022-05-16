package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/brackbk/purchase-service/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("dsn"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	/* 	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	} */

	if os.Getenv("AutoMigrateDb") == "true" {

		db.AutoMigrate(&domain.Purchase{}, &domain.PurchaseItem{})
		os.Setenv("AutoMigrateDb", "false")
	}

	return db.Debug()
}
