package main

import "fmt"

func main(){

	var mes int 

	fmt.Print("Ingrese el número del mes que desea saber: ")
	fmt.Scanf("%v", &mes)

	if mes > 0 && mes < 13 {
		switch {
			case mes == 1: 
				fmt.Println(mes,",Enero")
			case mes == 2: 
				fmt.Println(mes,",Febrero")
			case mes == 3: 
				fmt.Println(mes,",Marzo")
			case mes == 4: 
				fmt.Println(mes,",Abril")
			case mes == 5: 
				fmt.Println(mes,",Mayo")
			case mes == 6: 
				fmt.Println(mes,",Junio")
			case mes == 7: 
				fmt.Println(mes,",Julio")
			case mes == 8: 
				fmt.Println(mes,",Agosto")
			case mes == 9: 
				fmt.Println(mes,",Septiembre")
			case mes == 10: 
				fmt.Println(mes,",Octubre")
			case mes == 11: 
				fmt.Println(mes,",Noviembre")
			case mes == 12: 
				fmt.Println(mes,",Diciembre")
		}	
	}else {
		fmt.Println("El número ingresado",mes,"no pertenece a ningún mes.")
	}
}