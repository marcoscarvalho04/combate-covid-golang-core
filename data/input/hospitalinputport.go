package input

type HospitalInputPort struct {
	Nome        string `json:"nome"`
	Cnpj        string `json:"cnpj"`
	Endereco    string `json:"endereco"`
	Localizacao struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}
	PercentualDeOcupacao float32 `json:"percentualDeOcupacao"`
}
