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
}

func ConverterInputPortToOutputPort(hospitalInputPort input.HospitalInputPort) HospitalOutputPort {
	hospitalOutputPort := HospitalOutputPort{}

	hospitalOutputPort.Nome = hospitalInputPort.Nome
	hospitalOutputPort.Cnpj = hospitalInputPort.Cnpj
	hospitalOutputPort.Endereco = hospitalInputPort.Endereco
	hospitalOutputPort.Localizacao = hospitalInputPort.Localizacao
	hospitalOutputPort.PercentualDeOcupacao = hospitalInputPort.PercentualDeOcupacao

	return hospitalOutputPort
}
