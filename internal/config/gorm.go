package config

import (
	"log"
	"soulsaveronline-backend-go/internal/entity"

	"github.com/spf13/viper"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(v *viper.Viper) *gorm.DB {
	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "libsql",
		DSN:        v.GetString("DB_URL"),
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("gorm.Open: %v", err)
	}

	db.AutoMigrate(&entity.Device{})

	return db
}
