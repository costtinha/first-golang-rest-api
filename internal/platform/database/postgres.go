package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	appcfg "github.com/costtinha/first-golang-rest-api/internal/config"
	applg "github.com/costtinha/first-golang-rest-api/internal/logger"
)

func Connect(cfg *appcfg.Config, lg *applg.Logger) (*gorm.DB, error) {
	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}
	var db *gorm.DB
	var err error
	retries := 5

	for i := 0; i < retries; i++ {
		db, err = gorm.Open(postgres.Open(cfg.DSN()), gormCfg)
		if err == nil {
			break
		}
		lg.Warnf("Failed to connect to database, retrying (%d,%d): %v", i+1, retries, err)
		time.Sleep(2 * time.Second)

	}

	if err != nil {
		lg.Errorf("failed to connect to database after %d retries: %v", retries, err)
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	return db, nil
}
