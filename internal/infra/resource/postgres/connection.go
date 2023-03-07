package postgres

import (
	"fmt"

	"enerBit-system/config"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pgOptions struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
}

func (p *pgOptions) getDns() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.user, p.password, p.dbName)
}

func NewConnection() *gorm.DB {
	dns := pgOptions{
		host:     config.Environments().DbHost,
		port:     config.Environments().DbPort,
		user:     config.Environments().DbUser,
		password: config.Environments().DbPassword,
		dbName:   config.Environments().DbName,
	}

	dbInstance, err := gorm.Open(postgres.Open(dns.getDns()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	log.Info("Postgres Connection Successfully")
	return dbInstance
}
