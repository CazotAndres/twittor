package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword, rutina para encriptar el pwassword*/
func EncriptarPassword(pass string) (string, error) {
	//Esta es la cantidad de veces al cuadrado que va a encryptar la password
	costo := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
