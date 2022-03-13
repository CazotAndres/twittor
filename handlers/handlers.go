package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/CazotAndres/twittor/middlew"
	"github.com/CazotAndres/twittor/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores va entrar acá cunado llamemos a la API
func Manejadores() {
	router := mux.NewRouter()

	//router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")
	//Abrir el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//Permisos a cualquier con allowAll
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
