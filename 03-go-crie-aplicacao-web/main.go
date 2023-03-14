package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "Notebook", Descricao: "Vendido como novo", Valor: 150, Quantidade: 1},
		{"Bola de futebol", "Bola Nova", 15, 16},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
