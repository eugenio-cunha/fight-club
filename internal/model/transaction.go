package model

import "time"

type Transaction struct {
	Amount      int    `json:"valor"`
	Kind        string `json:"tipo"`
	Description string `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}
