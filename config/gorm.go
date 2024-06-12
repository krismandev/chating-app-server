package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	username := Config.DBUsername
	password := Config.DBPassword
	host := Config.DBHost
	port := Config.DBPort
	database := Config.DBName
	// idleConnection := viper.GetInt("database.pool.idle")
	// maxConnection := viper.GetInt("database.pool.max")
	// maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}

	_, err = db.DB()
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}

	// connection.SetMaxIdleConns(idleConnection)
	// connection.SetMaxOpenConns(maxConnection)
	// connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

// type logrusWriter struct {
// 	Logger *logrus.Logger
// }

// func (l *logrusWriter) Printf(message string, args ...interface{}) {
// 	l.Logger.Tracef(message, args...)
// }
