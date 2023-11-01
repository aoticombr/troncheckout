package main

type Departamento struct {
	ID        string `json:"id"`
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

type Cargo struct {
	ID        string `json:"id"`
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

type Contrato struct {
	CodigoDepartamento string       `json:"codigoDepartamento"`
	Departamento       Departamento `json:"departamento"`
	CodigoCargo        string       `json:"codigoCargo"`
	Cargo              Cargo        `json:"cargo"`
}

type Pessoa struct {
	ID        string   `json:"id"`
	Tipo      int      `json:"tipo"`
	Matricula string   `json:"matricula"`
	Situacao  int      `json:"situacao"`
	Inscricao string   `json:"inscricao"`
	Nome      string   `json:"nome"`
	PessoaID  string   `json:"pessoaId"`
	PisPasep  string   `json:"pisPasep"`
	Contrato  Contrato `json:"contrato"`
}
type Pessoas []Pessoa

// func main() {
// 	// Exemplo de como criar uma instância da estrutura e preenchê-la com os dados fornecidos
// 	pessoa := Pessoa{
// 		ID:        "1bfb80ab-d2d8-41bb-951c-216b4b780d86",
// 		Tipo:      1,
// 		Matricula: "239",
// 		Situacao:  1,
// 		Inscricao: "00583879195",
// 		Nome:      "PAULO HENRIQUE TADEU ALVES DE OLIVEIRA",
// 		PessoaID:  "a0348dae-95b2-4896-8dc7-33d93aa57564",
// 		PisPasep:  "20695019206",
// 		Contrato: Contrato{
// 			CodigoDepartamento: "15",
// 			Departamento: Departamento{
// 				ID:        "d45fff72-e81c-4771-90f1-cb6b140eccbc",
// 				Codigo:    "15",
// 				Descricao: "DESENV - FÁBRICA CARGAS",
// 			},
// 			CodigoCargo: "2",
// 			Cargo: Cargo{
// 				ID:        "51d59670-7cb0-420a-942c-f51cb8c2dc98",
// 				Codigo:    "2",
// 				Descricao: "ANALISTA DE SISTEMAS",
// 			},
// 		},
// 	}

// 	fmt.Printf("%+v\n", pessoa)
// }
