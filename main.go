package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"time"

	excel "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aoticombr/golang/http"
)

const (
	urllogin     = "https://apilogin.tron.com.br/api/autenticacao/login" //post
	urldadosuser = "https://apiconnect.tron.com.br/api/empregados"       //get
	urlponto     = "https://apiconnect.tron.com.br/api/frequencias"      //get
)

var (
	FUsers Users
	//Http     *http.THttp
	FToken   *TokenData
	FFiltros *Filtros
	FPessoa  *Pessoa
	FPontos  *Pontos
)

func pontos(ps *Pessoa, tk *TokenData, dtaini time.Time, dtafim time.Time) (*Pontos, error) {
	Http := http.NewHttp()
	Http.SetUrl(urlponto)
	Http.Request.Body = nil
	Http.Metodo = http.M_GET
	Http.EncType = http.ET_RAW
	Http.Request.Header.ContentType = "application/json"
	Http.Request.Header.Accept = "application/json"
	Http.Request.Header.ExtraFields.Add("UserAgent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	Http.Request.Header.ExtraFields.Add("serviceidentify", tk.Tenants[0].TenantId)

	Http.AuthorizationType = http.AT_Bearer
	Http.Authorization = "Bearer " + tk.Token

	filtros := NewFiltros()
	filtros.Add("inscricao", ps.Inscricao, "")
	filtros.Add("data", dtaini.Format("20060102"), "gte")
	filtros.Add("data", dtafim.Format("20060102"), "lte")

	Http.Params.Add("filtro", filtros.GetJsonString())

	Resp, Err := Http.Send()
	if Err != nil {
		return nil, fmt.Errorf("Erro ao enviar requisição: %v", Err)
	}
	fmt.Println("Status:", Resp.StatusCode)
	if Resp.StatusCode == 200 {
		var P *Pontos
		Err = json.Unmarshal(Resp.Body, &P)
		if Err != nil {
			return nil, fmt.Errorf("Erro ao enviar requisição: %v", Err)
		}
		return P, nil

	} else {
		fmt.Println("Status:", Resp.StatusCode)
		fmt.Println("Body:", string(Resp.Body))
		return nil, fmt.Errorf("Status Code: %v, Status Msg %v", Resp.StatusCode, Resp.StatusMessage)
	}

}
func pessoa(tk *TokenData) (*Pessoa, error) {
	Http := http.NewHttp()
	Http.SetUrl(urldadosuser)
	Http.Request.Body = nil
	Http.Metodo = http.M_GET
	Http.EncType = http.ET_RAW
	Http.Request.Header.ContentType = "application/json"
	Http.Request.Header.Accept = "application/json"
	Http.Request.Header.ExtraFields.Add("UserAgent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	Http.Request.Header.ExtraFields.Add("serviceidentify", tk.Tenants[0].TenantId)
	Http.AuthorizationType = http.AT_Bearer
	Http.Authorization = "Bearer " + tk.Token

	Resp, Err := Http.Send()
	if Err != nil {
		return nil, fmt.Errorf("Erro ao enviar requisição: %v", Err)
	}
	if Resp.StatusCode == 200 {
		var P *Pessoas
		err := json.Unmarshal(Resp.Body, &P)
		if err != nil {
			return nil, fmt.Errorf("Erro ao abrir o arquivo JSON:", err)

		}
		if len(*P) > 0 {
			return &(*P)[0], nil
		}
		return nil, fmt.Errorf("Dado não Encontrado")

	} else {
		fmt.Println("Status:", Resp.StatusCode)
		fmt.Println("Body:", string(Resp.Body))
		return nil, fmt.Errorf("Status Code: %v, Status Msg %v", Resp.StatusCode, Resp.StatusMessage)
	}
}
func login(u User) (*TokenData, error) {
	jsonbody, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}
	Http := http.NewHttp()
	Http.SetUrl(urllogin)
	Http.Request.Body = jsonbody
	Http.Metodo = http.M_POST
	Http.EncType = http.ET_RAW
	Http.Request.Header.ContentType = "application/json"
	Http.Request.Header.Accept = "application/json"
	Http.Request.Header.ExtraFields.Add("UserAgent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	Resp, Err := Http.Send()
	if Err != nil {
		return nil, fmt.Errorf("Erro ao enviar requisição: %v", Err)
	}
	if Resp.StatusCode == 200 {
		var T *TokenData
		json.Unmarshal(Resp.Body, &T)
		return T, nil

	} else {
		fmt.Println("Status:", Resp.StatusCode)
		fmt.Println("Body:", string(Resp.Body))
		return nil, fmt.Errorf("Status Code: %v, Status Msg %v", Resp.StatusCode, Resp.StatusMessage)
	}
}

func decimalToHoursMinutes(decimalValue float64) string {
	// Extrair a parte inteira (horas) do valor decimal
	hours := int(math.Floor(decimalValue))

	// Calcular a parte fracionária (minutos)
	minutes := int((decimalValue - float64(hours)) * 60)

	// Formatar o resultado como "HH:MM"
	result := fmt.Sprintf("%02d:%02d:00", hours, minutes)

	return result
}

func main() {
	// 1. Abrir o arquivo JSON
	arquivoJSON, err := ioutil.ReadFile("logins.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo JSON:", err)
		return
	}

	// 2. Decodificar o conteúdo JSON em uma estrutura de dados Go

	if err := json.Unmarshal(arquivoJSON, &FUsers); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	// Agora você pode acessar os dados na estrutura de dados 'usuarios'

	for _, u := range FUsers {
		FToken, err := login(u)
		if err != nil {
			fmt.Println("Erro ao fazer login:", err)
			return
		}
		fmt.Println("TenantId:", FToken.Tenants[0].TenantId)
		/*Captura dados importantes como CPF(inscrição)*/
		FPessoa, err := pessoa(FToken)
		if err != nil {
			fmt.Println("Erro ao fazer obter dados de pessoa:", err)
			return
		}
		//fmt.Println("Nome:", FPessoa)
		primeirodiadoano := time.Date(time.Now().Year()-1, 1, 1, 0, 0, 0, 0, time.UTC)
		FPontos, err = pontos(FPessoa, FToken, primeirodiadoano, time.Now())
		if err != nil {
			fmt.Println("Erro ao fazer obter dados de pessoa:", err)
			return
		}
		//fmt.Println("Pontos:", FPontos)
		f := excel.NewFile()
		SheetName := "Pontos"
		index := f.NewSheet(SheetName)
		f.SetActiveSheet(index)
		//
		f.SetCellValue(SheetName, "A1", "Data")
		f.SetCellValue(SheetName, "B1", "Hr. Ini")
		f.SetCellValue(SheetName, "C1", "Hr. Fim")
		f.SetCellValue(SheetName, "D1", "Total")
		f.SetCellValue(SheetName, "E1", "Total do dia")
		f.SetCellValue(SheetName, "F1", "Hora do dia")
		f.SetCellValue(SheetName, "G1", "Diferença")
		f.SetCellValue(SheetName, "H1", "Hora Decimal")
		f.SetCellValue(SheetName, "I1", "Credito Decimal")
		f.SetCellValue(SheetName, "J1", "Debito Decimal")
		f.SetCellValue(SheetName, "K1", "Observação")
		linha := 1
		var credito, debito float64 = 0.0, 0.0
		for _, p := range *FPontos {
			if len(p.Registros) > 0 {
				linha++
				linhastr := fmt.Sprintf("%d", linha)
				dataStr := fmt.Sprint(p.Data)
				data, err := time.Parse("20060102", dataStr)
				if err != nil {
					fmt.Println("Erro ao fazer o parsing da data:", err)
					return
				}
				dataFormatada := data.Format("02/01/2006")
				var (
					horarioInicial time.Time
					horarioFinal   time.Time
				)
				//	linhaP, linhaU := "", ""
				TotalHoras := 0.0
				var Horarios []time.Time
				for _, r := range p.Registros {
					horario, _ := time.Parse("15:04:05", r.Horario)
					Horarios = append(Horarios, horario)
				}
				sort.Slice(Horarios, func(i, j int) bool {
					return Horarios[i].Before(Horarios[j])
				})

				for index, hora := range Horarios {
					f.SetCellValue(SheetName, "A"+linhastr, dataFormatada)
					mod := math.Mod(float64(index), 2)
					if mod == 1 {
						horarioFinal = hora
						f.SetCellValue(SheetName, "C"+linhastr, hora.Format("15:04:05"))
						tempoTrabalhado := horarioFinal.Sub(horarioInicial)
						horas := tempoTrabalhado.Hours()
						TotalHoras += horas
						horas2 := decimalToHoursMinutes(horas)
						f.SetCellValue(SheetName, "D"+linhastr, horas2)

						linha++
						linhastr = fmt.Sprintf("%d", linha)
					} else {
						horarioInicial = hora
						f.SetCellValue(SheetName, "B"+linhastr, hora.Format("15:04:05"))
					}
				}
				TotalHorasSTR := decimalToHoursMinutes(TotalHoras)
				f.SetCellValue(SheetName, "E"+linhastr, TotalHorasSTR)
				TotalHorasTrabSTR := "08:48:00"
				f.SetCellValue(SheetName, "F"+linhastr, TotalHorasTrabSTR)
				H1, _ := time.Parse("15:04:05", TotalHorasSTR)
				H2, _ := time.Parse("15:04:05", TotalHorasTrabSTR)
				var tempoTrabalhado1 time.Duration
				NEGRITO1, _ := f.NewStyle(`{"font":{"bold":true,"color": "#0220FE"}}`)
				NEGRITO2, _ := f.NewStyle(`{"font":{"bold":true,"color": "#FF0000"}}`)
				NEGRITO3, _ := f.NewStyle(`{"font":{"bold":true}}`)
				ft := 0
				if H1.After(H2) {
					tempoTrabalhado1 = H1.Sub(H2)
					ft = 1
				} else if H1.Equal(H2) {
					ft = 3
					tempoTrabalhado1 = H1.Sub(H2)
				} else {
					ft = 2
					tempoTrabalhado1 = H2.Sub(H1)
				}
				tempoTrabalhado2 := tempoTrabalhado1.Hours()
				tempoTrabalhado3 := decimalToHoursMinutes(tempoTrabalhado2)
				f.SetCellValue(SheetName, "G"+linhastr, tempoTrabalhado3)
				switch ft {
				case 1:
					f.SetCellStyle(SheetName, "G"+linhastr, "G"+linhastr, NEGRITO1)
				case 2:
					f.SetCellStyle(SheetName, "G"+linhastr, "G"+linhastr, NEGRITO2)
				case 3:
					f.SetCellStyle(SheetName, "G"+linhastr, "G"+linhastr, NEGRITO3)
				}
				switch ft {
				case 1:
					credito += tempoTrabalhado1.Hours()
					f.SetCellValue(SheetName, "H"+linhastr, tempoTrabalhado1.Hours())
				case 2:
					debito += tempoTrabalhado1.Hours()
					f.SetCellValue(SheetName, "H"+linhastr, -tempoTrabalhado1.Hours())
				case 3:
					credito += tempoTrabalhado1.Hours()
					f.SetCellValue(SheetName, "H"+linhastr, tempoTrabalhado1.Hours())
				}
				f.SetCellValue(SheetName, "I"+linhastr, credito)
				f.SetCellValue(SheetName, "J"+linhastr, debito)
				f.SetCellValue(SheetName, "K"+linhastr, p.Observacao)

				linha++
			}
		}
		f.SaveAs("Pontos.xlsx")
	}

}
