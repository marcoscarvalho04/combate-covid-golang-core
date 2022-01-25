package input

import (
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/input"
	"github.com/marcoscarvalho04/combate-covid-golang-core/data/output"
	ports "github.com/marcoscarvalho04/combate-covid-golang-core/ports/output"
)

func SalvarHospital(inputPort input.HospitalInputPort) (output.HospitalOutputPort, error) {
	output := output.HospitalOutputPort{}
	outputPortData := output.ConverterInputPortToOutputPort(inputPort)
	return ports.SalvarHospital(outputPortData)

}
