package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBconnection(url string) *gorm.DB {
	newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: 200 * time.Millisecond,
            LogLevel:      logger.Warn,
            Colorful:      true,
        },
    )

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{ Logger: newLogger })
	if err != nil {
		log.Fatal(err)
	}
	
	return db
}