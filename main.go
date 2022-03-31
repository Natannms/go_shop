package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectDatabase()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func conexaoComBancoDeDados() {
	panic("unimplemented")
}

func index(w http.ResponseWriter, r *http.Request) {

	// db := conexaoComBancoDeDados()
	// rows, err := db.Query("SELECT * FROM alura_go.products")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// p := Products{}
	// produtos := []Products{}

	// for rows.Next() {
	// 	var ProductID int
	// 	var Name string
	// 	var Price float64
	// 	var Quantity int

	// 	err = rows.Scan(&ProductID, &Name, &Price, &Quantity)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	p.ProductID = ProductID
	// 	p.Name = Name
	// 	p.Price = Price
	// 	p.Quantity = Quantity

	// 	produtos = append(produtos, p)
	// }

	// fmt.Sprintln(rows)
	templates.ExecuteTemplate(w, "Index", "produtos")
}
