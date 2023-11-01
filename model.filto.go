package main

import "encoding/json"

type Filtro struct {
	Campo    string `json:"campo"`
	Valor    string `json:"valor"`
	Condicao string `json:"condicao"`
}

type Filtros []Filtro

func (f *Filtros) Add(campo string, valor string, condicao string) {
	*f = append(*f, Filtro{campo, valor, condicao})
}
func (f *Filtros) GetJsonString() string {
	jsonbody, err := json.Marshal(*f)
	if err != nil {
		return ""
	}
	return string(jsonbody)
}

func NewFiltros() *Filtros {
	return &Filtros{}
}
