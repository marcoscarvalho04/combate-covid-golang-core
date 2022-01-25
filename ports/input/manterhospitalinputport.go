package input

import (
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/input"
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/output"
)

func SalvarHospital(inputPort input.HospitalInputPort) (output.HospitalOutputPort, error) {
	output := output.HospitalOutputPort{}
	outputPortData := output.ConverterInputPortToOutputPort(inputPort)
	return output.Salvar.SalvarHospital(outputPortData)

}
