package config

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(
	dialector gorm.Dialector,
) *gorm.DB {
	dbClient, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	return dbClient
}
