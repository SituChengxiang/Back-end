package database

import (
	"Back-end/config"
	"fmt"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	user := config.Config.GetString("database.user")
	pass := config.Config.GetString("database.pass")
	port := config.Config.GetString("database.port")
	host := config.Config.GetString("database.host")
	name := config.Config.GetString("database.name")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Err(err).Msg("Database connect failed")
	}

	// 自动建表（已有不会覆盖）
	err = autoMigrate(db)
	if err != nil {
		log.Error().Err(err).Msg("Database migrate failed")
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
