package main

import (
	"fmt"
)

func Operacion(opc int, num1 int, num2 int) int {
	switch opc {
	case 1:
		return num1 + num2
	case 2:
		return num1 - num2
	case 3:
		return num1 * num2
	case 4:
		if num2 != 0 {
			return num1 / num2
		} else {
			fmt.Println("Error: No se puede dividir por cero.")
			return 0
		}
	case 5:
		return num1 * num2 / 100
	default:
		fmt.Println("Error: Operación no válida.")
		return 0
	}
}

func main() {
	for {
		var opcion int
		fmt.Println(
			"¡Hola bienvenido a calcugo! Ingrese que operación desea realizar: \n",
			"1. Sumar \n",
			"2. Restar \n",
			"3. Multiplicar \n",
			"4. Dividir \n",
			"5. Porcentaje \n",
			"6. Salir")
		fmt.Println("Ingrese la opción: ")
		fmt.Scanln(&opcion)

		if opcion == 6 {
			fmt.Println("Saliendo del programa... ")
			return
		}

		var num1, num2 int
		fmt.Println("Por favor ingrese el primer número: ")
		fmt.Scanln(&num1)
		fmt.Println("Por favor ingrese el segundo número: ")
		fmt.Scanln(&num2)

		resultado := Operacion(opcion, num1, num2)
		fmt.Printf("El resultado es: %v\n", resultado)
	}
}
