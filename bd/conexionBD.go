package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexion a la BD
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://andres:teto0412@cluster0.aisxv.mongodb.net/cluster0?retryWrites=true&w=majority")

//conectarBD es la funcion para conectar la BD
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion Exitosa con la BD")
	return client
}

//ChequeoConnection es el ping a la BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
