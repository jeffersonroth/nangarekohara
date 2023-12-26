package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func NewApplication() Application {
	app := Application{}
	app.Env = NewEnv()
	app.DB = Init(app.Env)
	return app
}
