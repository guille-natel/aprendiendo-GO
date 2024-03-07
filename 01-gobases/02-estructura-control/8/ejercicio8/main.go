package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	//Obtener edad de benjamin
	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])

	//Contar la cantidad de personas mayores 
	var contadorMayores int = 0

	for _, element := range employees {
		if(element > 21){
			contadorMayores++
		}
	}

	fmt.Println("La cantidad de empleados mayores son: ", contadorMayores)

	//Agregar un elemento al mapa
	fmt.Println(employees)
	employees["Federico"] = 25
	fmt.Println(employees)

	//Eliminar un elemento del mapa 
	delete(employees,"Pedro")
	fmt.Println(employees)


}
