package config

import (
	"fmt"
	"os"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql" // mysql driver

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent"
)

// MapConfig ...
type MapConfig struct {
	Database *Database
}

// Database ...
type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// NewMapConfig ...
func NewMapConfig() *MapConfig {
	database := &Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return &MapConfig{
		Database: database,
	}
}

// DBConnection ...
func (mc *MapConfig) DBConnection() (*ent.Client, error) {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		mc.Database.Username,
		mc.Database.Password,
		mc.Database.Host,
		mc.Database.Port,
		mc.Database.DBName)

	drv, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv)), nil
}
