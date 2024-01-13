package database

import (
	"fmt"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}

func NewDatabase(conf DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		conf.Host, conf.User, conf.Password, conf.DbName, conf.Port)

	return gorm.Open(postgres.New(postgres.Config{
		DriverName:           "nrpostgres",
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
}
