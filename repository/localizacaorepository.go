package repository

import "github.com/marcoscarvalho04/combate-covid-golang-core/data/entity"

type LocalizacaoRepository struct{}

func (c LocalizacaoRepository) SaveLocalizacao(newLocalizacao *entity.LocalizacaoEntity) (ID uint, err error) {
	database, err := ConfigureDatabase()
	if err != nil {
		ID = 0
		return ID, err
	}
	responseMigration := database.Db.AutoMigrate(newLocalizacao)
	if responseMigration != nil {
		err = responseMigration
		ID = 0
		return ID, err
	}
	responseCreation := database.Db.Create(newLocalizacao)
	if responseCreation.Error != nil {
		err = responseCreation.Error
		ID = 0
		return ID, err
	}
	ID = newLocalizacao.ID
	err = nil
	return ID, err
}
