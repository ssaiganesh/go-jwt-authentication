package database

import (
	sql "github.com/go-sql-driver/mysql"
	"github.com/ssaiganesh/go-jwt-authentication/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := sql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "testdb",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	connection, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(models.User{})
}
