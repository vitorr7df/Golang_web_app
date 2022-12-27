package main

import (
	"net/http"

	"github.com/vitorr7df/routes"
)

func main() { //CRIAR SERVIDOR E EXIBIR NUMA PAGINA HTML
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
