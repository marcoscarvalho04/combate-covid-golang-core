package database

import (
	"github.com/marcoscarvalho04/combate-covid-golang-core/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	PathConfig string
	Db         *gorm.DB
	Config     configuration.Config
}

func (a *Database) LoadDatabase() error {
	if a.Db != nil {
		return nil
	}

	a.Config = configuration.LoadAllConfiguration(a.PathConfig)

	db, err := gorm.Open(mysql.Open(a.Config.Database.Url), &gorm.Config{})

	if err != nil {
		return err
	}
	a.Db = db

	return nil

}
