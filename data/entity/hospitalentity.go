package entity

import (
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/output"
	"gorm.io/gorm"
)

type LocalizacaoEntity struct {
	Id        uint `gorm:"primaryKey"`
	Latitude  string
	Longitude string
}
type HospitalEntity struct {
	gorm.Model
	Id                   uint `gorm:"primaryKey"`
	Nome                 string
	Cnpj                 string
	Endereco             string
	Localizacao          LocalizacaoEntity
	PercentualDeOcupacao float32
}

func ConverterOutputToEntity(outputPort output.HospitalOutputPort) HospitalEntity {
	return HospitalEntity{
		Id:                   uint(outputPort.Id),
		Nome:                 outputPort.Nome,
		Cnpj:                 outputPort.Cnpj,
		Endereco:             outputPort.Cnpj,
		PercentualDeOcupacao: outputPort.PercentualDeOcupacao,
		Localizacao: LocalizacaoEntity{
			Latitude:  outputPort.Localizacao.Latitude,
			Longitude: outputPort.Localizacao.Longitude,
		},
	}
}
