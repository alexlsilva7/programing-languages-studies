package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	println("Servidor iniciado en http://localhost:8090")
	http.HandleFunc("/", index)
	// Iniciar servidor web
	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Camiseta Polo", Preco: 79.90, Quantidade: 10},
		{Nome: "Notebook", Descricao: "Notebook Dell", Preco: 1999.90, Quantidade: 5},
		{Nome: "Caneta", Descricao: "Caneta Bic", Preco: 1.90, Quantidade: 100},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
