package main

import (
	"log"

	"github.com/CazotAndres/twittor/bd"
	"github.com/CazotAndres/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexión a la BD")
		return
	}
	handlers.Manejadores()
}
