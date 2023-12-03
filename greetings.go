package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada
func Hello(nombre string) (string, error) {

	if nombre == "" {
		return nombre, errors.New("Nombre vacío")
	}

	// Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), nombre)
	return message, nil
}

//Hellos devuelve un mapa que está asociado a cada uno de los nombres con su mensaje
func Hellos(nombres []string)(map[string]string, error) {
	messages := make(map[string]string)

	for _, nombre := range nombres {
		message, err := Hello(nombre)
		if err != nil {
			return nil, err
		}

		messages[nombre] = message
	}
	
	return messages, nil
}

// randomFormat devuelve uno de varios formatos de mensajes de saludos
// Se selecciona de forma aleatoria
func randomFormat() string {
	formats := [] string {
		"Hola, %v! Bienvenido",
		"Que bueno verte, %v",
		"Saludos amigo %v",
	}

	return formats[rand.Intn(len(formats))]
}