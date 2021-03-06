package controllers

import (
	"github.com/David-Nascimento/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))


func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request)  {
	temp.ExecuteTemplate(w,"New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//Converte o preco para float64
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco: ", err)
		}

		//Converte a quantidade para int
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade: ", err)
		}

		models.CriarNovoProduto(
			nome,
			descricao,
			precoConvertido,
			quantidadeConvertida,
		)
	}
	http.Redirect(w,r,"/", 301)
}