package main

import (
	"html/template"
	"net/http"
	"github.com/LuqAndrade/db"
	_ "github.com/lib/pq"
)

type Produto struct {
	id int;
	Nome string;
	Descricao string;
	Preco float64;
	Quantidade int;
}


var temp = template.Must(template.ParseGlob("templates/*.html"))

func main(){
	
	http.HandleFunc("/", index)
	http.ListenAndServe(":5500",nil)
	
}
func index(w http.ResponseWriter,r *http.Request){
	db := db.Conectacombancodedados()

	selectdetodososprodutos, err := db.Query("select * from produtos")
	if err != nil{
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectdetodososprodutos.Next(){
		var id int
		var nome,descricao string
		var preco float64
		var quantidade int

		err = selectdetodososprodutos.Scan(&id,&nome,&descricao,&preco,&quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos,p)
		
	}
	temp.ExecuteTemplate(w,"index",produtos)
	defer db.Close()
}


