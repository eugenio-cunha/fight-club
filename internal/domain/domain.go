package domain

import (
	"fight-club/internal/database"
	"fight-club/internal/model"

	"github.com/bytedance/sonic"
)

func Balance(id int64) (*[]byte, error) {
	var buf []byte
	row := database.Row("SELECT process_balance($1)", id)
	if err := row.Scan(&buf); err != nil {
		return nil, err
	}

	return &buf, nil
}

func Transaction(id int64, value int, kind string, description string) (*[]byte, error) {
	var limit, balance int

	row := database.Row("CALL process_transaction($1, $2, $3, $4, $5 ,$6)", id, kind, value, description, 0, 0)
	if err := row.Scan(&limit, &balance); err != nil {
		return nil, err
	}

	buf, err := sonic.Marshal(model.Statement{
		Limit:   limit,
		Balance: balance,
	})
	if err != nil {
		return nil, err
	}

	return &buf, nil
}