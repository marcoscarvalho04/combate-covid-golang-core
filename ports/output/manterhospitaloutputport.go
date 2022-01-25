package output

import (
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/entity"
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/output"
	"github.com/marcoscarvalho04/combate-covid-golang-core/repository"
)

func SalvarHospital(outputPort output.HospitalOutputPort) (output.HospitalOutputPort, error) {

	hospitalEntity := entity.ConverterOutputToEntity(outputPort)
	database := repository.Database{}
	err := database.Create(hospitalEntity)
	if err != nil {
		return outputPort, err
	}
	outputPort.Id = int32(hospitalEntity.ID)

	return outputPort, nil
}
