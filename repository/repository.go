package repository

import "github.com/marcoscarvalho04/combate-covid-golang-core/database"

func ConfigureDatabase() (database.Database, error) {
	database := database.Database{}
	err := database.LoadDatabase()
	return database, err
}
