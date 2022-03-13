package bd

import (
	"context"
	"time"

	"github.com/CazotAndres/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario chequea existencia en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//Este cancel, al estar con el defer se va ejecutar a lo ultimo de todo y le va dar de baja la timeout
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	//Va ir a la coleccion y solo va buscar un registro
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
