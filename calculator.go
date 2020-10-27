package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func leerOperador() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operador := scanner.Text()
	return operador
}

func LeerValores() [2]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	valores := scanner.Text()
	listaValores := strings.Split(valores, ",")
	valor1, error1 := strconv.Atoi(listaValores[0]) //Atoi parsea los valores a Integers
	valor2, error2 := strconv.Atoi(listaValores[1])
	if error1 != nil || error2 != nil {
		fmt.Println("Alguno de los valores no se pudo parsear.")
	}
	var listaValoresReturn [2]int //Atoi parsea los valores a Integers
	listaValoresReturn[0] = valor1
	listaValoresReturn[1] = valor2
	return listaValoresReturn
}

func suma(valor1 int, valor2 int) int {
	return valor1 + valor2
}

func resta(valor1 int, valor2 int) int {
	return valor1 - valor2
}

func multiplicacion(valor1 int, valor2 int) int {
	return valor1 * valor2
}

func division(valor1 int, valor2 int) int {
	return valor1 / valor2
}

func (calc) calcular(operador string, valor1 int, valor2 int) int {
	switch operador {
	case "+":
		resultado := suma(valor1, valor2)
		return resultado
	case "-":
		resultado := resta(valor1, valor2)
		return resultado
	case "*":
		resultado := multiplicacion(valor1, valor2)
		return resultado
	case "/":
		resultado := division(valor1, valor2)
		return resultado
	default:
		fmt.Printf("No se encontró un operador válido\n")
		return 0
	}
}
