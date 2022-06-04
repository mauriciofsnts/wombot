package database

import (
	"fmt"

	"code.db.cafe/wombot/internal/config"
	"code.db.cafe/wombot/internal/database/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB = new(gorm.DB)

func Start() error {
	cfg := config.Wombot.Pg

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=ETC/GMT", cfg.Host, cfg.Username, cfg.Password, cfg.DbName, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	// wtf ?
	*Db = *db

	return Db.AutoMigrate(&entities.User{}, &entities.Guild{}, &entities.Submission{})
}
