package database

import (
	"fmt"
	"undefeated-davout/echo-api-sample/config"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (db repositories.DBer, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword,
		cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
