package models
import "github.com/LuqAndrade/db"
type Produto struct {
	id int;
	Nome string;
	Descricao string;
	Preco float64;
	Quantidade int;
}

func buscatodososprodutos() []Produto{
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
	defer db.Close()
	return db
}