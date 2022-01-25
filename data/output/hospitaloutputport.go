package output

import "github.com/marcoscarvalho04/combate-covid-golang-core/data/input"

type HospitalOutputPort struct {
	Id          int32  `json:"id"`
	Nome        string `json:"nome"`
	Cnpj        string `json:"cnpj"`
	Endereco    string `json:"endereco"`
	Localizacao struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}
	PercentualDeOcupacao float32 `json:"percentualDeOcupacao"`

	Salvar SalvarHospital
}
type SalvarHospital interface {
	SalvarHospital(outputPort HospitalOutputPort) (HospitalOutputPort, error)
}

func (c *HospitalOutputPort) ConverterInputPortToOutputPort(hospitalInputPort input.HospitalInputPort) HospitalOutputPort {
	return HospitalOutputPort{
		Nome:                 hospitalInputPort.Nome,
		Cnpj:                 hospitalInputPort.Cnpj,
		Endereco:             hospitalInputPort.Endereco,
		PercentualDeOcupacao: hospitalInputPort.PercentualDeOcupacao,
		Localizacao:          hospitalInputPort.Localizacao,
	}
}

//
