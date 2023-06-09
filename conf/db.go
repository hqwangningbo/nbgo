package conf

import (
	"github.com/hqwangningbo/nbgo/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Error
	if !viper.GetBool("mode.develop") {
		logMode = logger.Info
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},

		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.User{})

	return db, nil
}
