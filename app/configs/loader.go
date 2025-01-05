package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Database Database
	Jwt      Jwt
}

type App struct {
	Name, Env, Version string
	Port               int
	Debug              bool
}

type Jwt struct {
	Key string
	Exp int
}

type Database struct {
	Host, Port, Name, User, Pass, Tz string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		panic("Error when load file configuration " + err.Error())
	}
	appDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		log.Fatal(err)
		appDebug = false
	}
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
		appPort = 8080
	}

	return &Config{
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Env:     os.Getenv("APP_ENV"),
			Port:    appPort,
			Debug:   appDebug,
			Version: os.Getenv("APP_VERSION"),
		},
		Jwt: Jwt{
			Key: os.Getenv("JWT_KEY"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Tz:   os.Getenv("DB_TZ"),
		},
	}

}
