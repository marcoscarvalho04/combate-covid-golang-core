package repository

import (
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/entity"
)

type HospitalRepository struct{}

func (c *HospitalRepository) Create(newHospital *entity.HospitalEntity) error {
	database, err := ConfigureDatabase()
	if err != nil {
		return err
	}
	localizacaoRepository := LocalizacaoRepository{}

	newHospital.LocalizacaoId, err = localizacaoRepository.SaveLocalizacao(&newHospital.LocalizacaoEntity)
	if err != nil {
		return err
	}
	resultMigration := database.Db.AutoMigrate(&newHospital)
	if resultMigration != nil {
		return resultMigration
	}

	resultCreate := database.Db.Create(&newHospital)

	if resultCreate.Error != nil {
		return resultCreate.Error
	}

	return nil
}
