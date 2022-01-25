package repository

import (
	"errors"

	"github.com/marcoscarvalho04/combate-covid-golang-core/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	Database struct {
		Url string `yaml:"url"`
	}
}

type Database struct {
	Db     *gorm.DB
	Config *config
}

func (a *Database) InitializeDatabase() error {
	if a.Db != nil {
		return nil
	}

	configuration := configuration.AllConfiguration{}
	a.Config = &config{}
	configuration.LoadAllConfiguration(a.Config)

	db, err := gorm.Open(mysql.Open(a.Config.Database.Url), &gorm.Config{})

	if err != nil {
		return err
	}
	a.Db = db

	return nil

}
func (a *Database) Create(model interface{}) error {
	if model == nil {
		return errors.New("model not defined")
	}
	a.InitializeDatabase()
	resultCreate := a.Db.Create(&model)

	if resultCreate.Error != nil {
		return errors.New(resultCreate.Error.Error())
	}
	return nil
}
