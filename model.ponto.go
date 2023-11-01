package main

type Registro struct {
	Horario         string `json:"horario"`
	HorarioCadastro string `json:"horarioCadastro"`
	Origem          int    `json:"origem"`
	DataFrequencia  int    `json:"dataFrequencia"`
	DescricaoOrigem string `json:"descricaoOrigem"`
	DadosBatida     struct {
		TipoBatida         int        `json:"tipoBatida"`
		InternetConnection bool       `json:"internetConnection"`
		Geolocalizacao     []struct{} `json:"geolocalizacao"`
		CodigoHash         string     `json:"codigoHash"`
		Nsr                int        `json:"nsr"`
		NsrFormatted       string     `json:"nsrFormatted"`
		Inpi               string     `json:"inpi"`
	} `json:"dadosBatida"`
}

type Item struct {
	ID                     string     `json:"id"`
	Inscricao              string     `json:"inscricao"`
	Data                   int        `json:"data"`
	TipoDia                int        `json:"tipoDia"`
	CargaHoraria           string     `json:"cargaHoraria"`
	Credito                string     `json:"credito"`
	Saldo                  string     `json:"saldo"`
	TipoSaldo              string     `json:"tipoSaldo"`
	ToleranciaEntradaSaida int        `json:"toleranciaEntradaSaida"`
	ToleranciaMaximaDiaria int        `json:"toleranciaMaximaDiaria"`
	HorasTrabalhadas       string     `json:"horasTrabalhadas"`
	Registros              []Registro `json:"registros"`
	Observacao             string     `json:"observacao"`
}

type Pontos []Item
