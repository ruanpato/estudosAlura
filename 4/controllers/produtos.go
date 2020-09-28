package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

// New ...
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

// Insert ...
func Insert(w http.ResponseWriter, r *http.Request) {
	var produto models.Produto
	if r.Method == "POST" {
		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")

		aux, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}
		produto.Preco = aux

		produto.Quantidade, err = strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		models.CriarNovoProduto(produto)
	}
	http.Redirect(w, r, "/", 301)
}

// Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

// Edit ...
func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.SelectProdutoByID(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

// Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	var produto models.Produto
	if r.Method == "POST" {
		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")

		aux, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float64: ", err)
		}
		produto.Preco = aux

		produto.Quantidade, err = strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da quantidade Para INT: ", err)
		}

		produto.ID, err = strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Erro na conversão da ID Para INT: ", err)
		}

		models.AtualizaProduto(produto)
	}
	http.Redirect(w, r, "/", 301)
}
