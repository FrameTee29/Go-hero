package bootstrap

import "gohero/pkg/database"

type Application struct {
	Env *Env
	DB  *database.Db
}

func App() Application {

	app := &Application{}

	app.Env = NewEnv()
	env := app.Env

	db, err := database.ConnectPostgresDB(
		database.WithHost(env.DbHost),
		database.WithPort(env.DbPort),
		database.WithUser(env.DbUsername),
		database.WithPassword(env.DbPassword),
		database.WithDbName(env.DbName),
		database.WithSSLMode(env.DbSslMode),
		database.WithTimezone(env.DbTimeZone),
	)
	if err != nil {
		panic("Can't connect database !!")
	}

	app.DB = db

	return *app
}
