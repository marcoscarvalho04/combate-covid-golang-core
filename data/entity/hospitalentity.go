package entity

import (
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/output"
	"gorm.io/gorm"
)

type LocalizacaoEntity struct {
	ID        uint `gorm:"primaryKey"`
	Latitude  string
	Longitude string
}
type HospitalEntity struct {
	gorm.Model
	ID                   uint `gorm:"primaryKey"`
	Nome                 string
	Cnpj                 string
	Endereco             string
	LocalizacaoEntity    LocalizacaoEntity `gorm:"-"`
	LocalizacaoId        uint              `gorm:"TYPE:integer REFERENCES LocalizacaoEntity;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PercentualDeOcupacao float32
}

func ConverterOutputToEntity(outputPort output.HospitalOutputPort) HospitalEntity {
	return HospitalEntity{

		Nome:                 outputPort.Nome,
		Cnpj:                 outputPort.Cnpj,
		Endereco:             outputPort.Cnpj,
		PercentualDeOcupacao: outputPort.PercentualDeOcupacao,
		LocalizacaoEntity: LocalizacaoEntity{
			Latitude:  outputPort.Localizacao.Latitude,
			Longitude: outputPort.Localizacao.Longitude,
		},
	}
}
