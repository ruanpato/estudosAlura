package models

import (
	"../db"
)

// Produto ...
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// BuscaTodosOsProdutos ...
func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

// CriarNovoProduto ...
func CriarNovoProduto(produto Produto) {
	db := db.ConectaComBancoDeDados()
	insereNoBancoDeDados, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)") // Postgres $i, mysql ?
	if err != nil {
		panic(err.Error())
	}
	insereNoBancoDeDados.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	defer db.Close()
}

// DeletaProduto ...
func DeletaProduto(idProduto string) {
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(idProduto)
	defer db.Close()
}

// SelectProdutoByID ...
func SelectProdutoByID(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()

	selectProduto, err := db.Query("SELECT * FROM produtos WHERE id=$1", idProduto)
	if err != nil {
		panic(err.Error())
	}
	produto := Produto{}
	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error)
		}
		produto.ID = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}
	defer db.Close()
	return produto
}

// AtualizaProduto ...
func AtualizaProduto(produto Produto) {
	db := db.ConectaComBancoDeDados()
	AtualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.ID)
	defer db.Close()
}
