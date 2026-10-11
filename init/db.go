package initializer

import (
	"api/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectToDatabase(cfg config.Config) (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	var err error
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}

func closeDatabase() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL db instance: %v", err)
	}

	return sqlDB.Close()
}
