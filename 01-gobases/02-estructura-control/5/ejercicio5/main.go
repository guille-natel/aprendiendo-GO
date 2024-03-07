package main

import "fmt"



func main() {

	var palabra string
	var contadorLetras int = 0

	fmt.Print("Ingrese una palabra: ")

	fmt.Scanf("%v", &palabra)

	fmt.Println("Palabra ingresada: ", palabra)

	for i := 0; i < len(palabra); i++ {
			fmt.Println("Letra:", palabra[i]) //arreglar 
			contadorLetras++
	}

	fmt.Println("Total de letras de la palabra: " , contadorLetras)


}
