package main

import (
	"net/http"

	"../4/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
