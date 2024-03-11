package bootstrap

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	// * Application
	AppEnv string `env:"APP_ENV"`
	Port   string `env:"PORT"`

	// * Database
	DbHost     string `env:"DB_HOST"`
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
	DbPort     int    `env:"DB_PORT"`
	DbSslMode  string `env:"DB_SSLMODE"`
	DbTimeZone string `env:"DB_TIMEZONE"`
}

func NewEnv() *Env {
	env := &Env{}

	env, err := loadEnv()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// * Default port
	if env.Port == "" {
		env.Port = "3000"
	}

	return env
}

func loadEnv() (*Env, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	env := &Env{}
	v := reflect.ValueOf(env).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Get the env tag value
		tag := fieldType.Tag.Get("env")
		if tag == "" {
			continue
		}

		t := fieldType.Type.Kind()

		envValue := os.Getenv(tag)

		// Check king of type
		switch t {
		// String
		case reflect.String:
			field.SetString(envValue)
			// Int
		case reflect.Int:
			intValue, err := strconv.Atoi(envValue)
			if err != nil {
				log.Fatal("Environment", envValue, "can't be loaded", err)
			}
			field.SetInt(int64(intValue))
			// default type is string
		default:
			field.SetString(envValue)
		}

	}

	return env, nil

}
