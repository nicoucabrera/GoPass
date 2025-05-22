package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Funcion para generar una contraseña aleatoria
// La contraseña tendrá una longitud de 12 caracteres, alternando entre letras y números

func newPass() string {
	rand.Seed(time.Now().UnixNano()) // Inicializa la semilla del generador de números aleatorios
	var letras = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numeros = "0123456789"
	var pass = make([]byte, 12) // Crea un slice de bytes para almacenar la contraseña
	// Alterna entre letras y números
	for i := 0; i < 12; i++ {
		// Si el índice es par, selecciona una letra, si es impar, selecciona un número
		if i%2 == 0 {
			pass[i] = letras[rand.Intn(len(letras))] // Selecciona una letra aleatoria
		} else if i%3 == 0 { // Si el índice es múltiplo de 3, selecciona una letra
			// Esto asegura que la contraseña tenga al menos una letra en posiciones impares
			pass[i] = letras[rand.Intn(len(letras))] // Selecciona una letra aleatoria
		} else {
			pass[i] = numeros[rand.Intn(len(numeros))] // Selecciona un número aleatorio
		}
	}
	return string(pass)
}

func passTxt(nombreTxt, contenido string) error {
	return os.WriteFile(nombreTxt, []byte(contenido), 0644) // Crea un archivo de texto con la contraseña generada
}

func main() {
	// Generar una contraseña aleatoria
	pass := newPass()
	fmt.Println("Generando contraseña...")
	time.Sleep(3 * time.Second)
	fmt.Println("Contraseña generada:", pass)
	fmt.Println("Generando archivo...")
	time.Sleep(3 * time.Second)
	err := passTxt("pass.txt", "Contraseña: "+pass)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
	} else {
		fmt.Println("Archivo creado con éxito")
	}
}
