package db

import (
	"database/sql"

	_ "github.com/lib/pq" //
)

// ConectaComBancoDeDados ...
func ConectaComBancoDeDados() *sql.DB {
	conexao := "host=localhost port=5432 user=db password=b1b2b3b4 dbname=alura_loja sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
