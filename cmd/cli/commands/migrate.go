package commands

import (
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getMigrate() (*migrate.Migrate, error) {
	dir := "file://./app/database/migrations"
	gormDB, err := configs.InitDB()

	if err != nil {
		return nil, err
	}

	db, err := gormDB.DB()

	if err != nil {
		return nil, err
	}

	driver, errDriver := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: "migrations",
	})

	if errDriver != nil {
		return nil, errDriver
	}

	m, errMigrate := migrate.NewWithDatabaseInstance(
		dir,
		"mysql",
		driver,
	)

	if errMigrate != nil {
		return nil, errMigrate
	}

	return m, nil
}
