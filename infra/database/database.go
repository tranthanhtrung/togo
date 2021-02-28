package database

import (
	"database/sql"
	"fmt"

	"github.com/manabie-com/togo/infra/config"
)

// NewPostgres create postgres database
func NewPostgres(host, user, password, dbname string, port int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	return sql.Open("postgres", psqlInfo)
}

// NewSQLLite create sqlLite database
func NewSQLLite(path string) (*sql.DB, error) {
	return sql.Open("sqlite3", path)
}

// NewDatabase create database for project
func NewDatabase(cfg config.Config) (*sql.DB, error) {
	switch cfg.Database.Type {
	case "postgres":
		return NewPostgres(
			cfg.Database.Host,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.DBName,
			cfg.Database.Port,
		)
	default:
		return NewSQLLite("./data.db")
	}
}
