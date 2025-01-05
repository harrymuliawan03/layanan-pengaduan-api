package configs

import (
	"fmt"
	"os"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/facades"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

type dbCon struct {
	username, password, db, host, port string
}

var (
	dbInstance *gorm.DB
	err        error
)

func InitDB() (db *gorm.DB, err error) {
	dbCon := dbCon{
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		db:       os.Getenv("DB_DATABASE"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbCon.username,
        dbCon.password,
        dbCon.host,
        dbCon.port,
        dbCon.db,
    )
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})


	if err != nil {
		panic(err)
	}

	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	return
}

func ConnectDB() {
	if dbInstance == nil {
		dbInstance, err = InitDB()
	}

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	facades.MakeOrm(dbInstance)
}
