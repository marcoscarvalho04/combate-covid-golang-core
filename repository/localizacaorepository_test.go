package repository

import (
	"testing"

	"github.com/marcoscarvalho04/combate-covid-golang-core/data/entity"
)

func TestCreateSuccessfullyModel(t *testing.T) {
	hospitalEntity := entity.HospitalEntity{
		Nome:     "teste",
		Cnpj:     "1234586789",
		Endereco: "teste",
		LocalizacaoEntity: entity.LocalizacaoEntity{
			Latitude:  "123",
			Longitude: "123",
		},
		PercentualDeOcupacao: 32.0,
	}
	hospitalRepository := HospitalRepository{}
	err := hospitalRepository.Create(hospitalEntity)

	if err != nil {
		t.Errorf("error creating new entity: %s", err.Error())
		return
	}
	t.Log("TestCreateSucessfullyModel passed!")
}
