package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CazotAndres/twittor/bd"
	"github.com/CazotAndres/twittor/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	//El body es Stream, una vez que se leyo, se destruye
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña al menos de 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya Existe un usuario con ese mail ", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
