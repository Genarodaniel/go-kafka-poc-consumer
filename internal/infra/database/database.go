package database

import (
	"database/sql"
	"fmt"
	"go-kafka-poc-consumer/config"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connect() *sql.DB {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBName,
		config.Config.DBHost,
		config.Config.DBPort,
		"disable",
	)

	fmt.Println(connectionString)

	db, err := sql.Open(config.Config.DBDriver, connectionString)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		panic(err)
	}

	return db
}

func Migrate(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file:..//internal/infra/database/migrations", config.Config.DBDriver, driver)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		fmt.Printf("DATABASE MIGRATIONS: %s", err.Error())
	}
}
