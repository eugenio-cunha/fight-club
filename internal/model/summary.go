package model

type Summary struct {
	Balance     Balance       `json:"saldo"`
	Transaction []Transaction `json:"ultimas_transacoes"`
}
