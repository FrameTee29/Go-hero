package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	DriverDb *gorm.DB
}

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	DdName   string
	Port     int
	SSLMode  string
	Timezone string
}

// ? Postgres option config

type Option func(*PostgresConfig)

func defaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "",
		User:     "",
		Password: "",
		DdName:   "",
		Port:     5432,
		SSLMode:  "",
		Timezone: "",
	}
}

func WithHost(host string) Option {
	return func(cfg *PostgresConfig) {
		cfg.Host = host
	}
}

func WithUser(user string) Option {
	return func(cfg *PostgresConfig) {
		cfg.User = user
	}
}

func WithPassword(password string) Option {
	return func(cfg *PostgresConfig) {
		cfg.Password = password
	}
}

func WithDbName(dbName string) Option {
	return func(cfg *PostgresConfig) {
		cfg.DdName = dbName
	}
}

func WithPort(port int) Option {
	return func(cfg *PostgresConfig) {
		cfg.Port = port
	}
}

func WithSSLMode(sslMode string) Option {
	return func(cfg *PostgresConfig) {
		cfg.SSLMode = sslMode
	}
}

func WithTimezone(timezone string) Option {
	return func(cfg *PostgresConfig) {
		cfg.Timezone = timezone
	}
}

func getDSN(cfg PostgresConfig) string {

	if cfg.Host == "" {
		panic("[Postgres Config] - Host is missing")
	}

	if cfg.User == "" {
		panic("[Postgres Config] - User is missing")
	}

	if cfg.Password == "" {
		panic("[Postgres Config] - Password is missing")
	}

	if cfg.SSLMode == "" {
		panic("[Postgres Config] - SSLMode is missing")
	}

	if cfg.Timezone == "" {
		panic("[Postgres Config] - Timezone is missing")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.DdName, cfg.Port, cfg.SSLMode, cfg.Timezone)

	return dsn
}

func ConnectPostgresDB(options ...Option) (*Db, error) {
	db := Db{}

	cfg := defaultPostgresConfig()

	for _, option := range options {
		option(&cfg)
	}

	dsn := getDSN(cfg)

	dial := postgres.Open(dsn)

	postgresCon, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.DriverDb = postgresCon

	return &db, nil
}
