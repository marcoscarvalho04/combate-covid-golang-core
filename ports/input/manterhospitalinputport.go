package input

import "github.com/marcoscarvalho04/combate-covid-golang-core/data/input"

type HospitalOperations interface {
	SalvarHospital(input.HospitalInputPort)
}

type HospitalInputPort struct {
	Operations HospitalOperations
}

func (c *HospitalInputPort) SalvarHospital(input input.HospitalInputPort) {

}
