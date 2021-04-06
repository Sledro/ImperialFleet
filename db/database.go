package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *sql.DB

// ConnectToDatabase - Connect to mysql db
func ConnectToDatabase() {
	var err error
	DB, err = sql.Open("mysql", viper.GetString("dbConnString"))
	if err != nil {
		log.Error("Could not connect to db: ", err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
