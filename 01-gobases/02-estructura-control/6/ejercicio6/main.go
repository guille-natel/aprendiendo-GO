package main

import "fmt"

func main(){

	var age int 
	var is_empleado string
	var antiguedad int
	var sueldo int 

	fmt.Println("Ingrese su edad: ")
	fmt.Scanf("%v", &age)

	fmt.Println("¿Usted esta empleado? S(Si) / N(No) ")
	fmt.Scanf("%v", &is_empleado)

	fmt.Println("Ingrese cuantos años de antiguedad tiene: ")
	fmt.Scanf("%v", &antiguedad)

	fmt.Println("Ingrese su sueldo: ")
	fmt.Scanf("%v", &sueldo)

	fmt.Println("Edad: ", age)
	fmt.Println("Empleado: ", is_empleado)
	fmt.Println("Antigüedad:", antiguedad)
	fmt.Println("Sueldo: ", sueldo)


	if age > 22 && is_empleado == "S" && antiguedad > 1 && sueldo > 100000 {
		fmt.Println("Felicitaciones! Usted cumple con los requisitos para el prestamo sin intereses.")
		} else if age > 22 && is_empleado == "S" && antiguedad > 1 && sueldo <= 100000 {
			fmt.Println("Felicitaciones! Usted cumple con los requisitos para el prestamo con intereses.") 
			} else {
				fmt.Println("Lo sentimos mucho! Usted no cumple con los requisitos para el prestamo.")
			}
					
}
	

	
// Para mejorar: mejorar que no ingresen numeros negativos, tema de minusculas y mayusculas para el string, que no ingresen 
// cadenas vacias
