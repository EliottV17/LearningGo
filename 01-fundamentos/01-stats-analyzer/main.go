package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
EJERCICIO 1.1: Analizador Estadístico de Métricas (CLI)

OBJETIVO:
Consolidar el control de flujo idiomático, funciones con retornos múltiples,
manejo explícito de errores y zero-values sin recurrir a excepciones.

REQUERIMIENTOS:
1. Definir una función `calculateStats(numbers []float64) (min float64, max float64, avg float64, err error)`:
   - Si el slice está vacío o es nil, debe retornar un error explícito creado con `errors.New` o `fmt.Errorf`.
   - Si contiene datos, iterar usando un único bucle `for ... range` para calcular mínimo, máximo y promedio.
   - Usar retornos nombrados (named return values) o retorno directo tradicional (evaluar legibilidad).

2. Implementar una función `parseInput(args []string) ([]float64, error)`:
   - Parsear los argumentos pasados por CLI (excluyendo el nombre del programa) usando `strconv.ParseFloat`.
   - Si algún argumento falla al parsearse, retornar el error envuelto indicando el argumento erróneo.

3. En `main()`:
   - Tomar los argumentos de `os.Args[1:]`.
   - Manejar los casos borde: si no hay argumentos, mostrar el mensaje de uso esperado y salir con código de salida no-cero (`os.Exit(1)`).
   - Aplicar `defer` para imprimir un mensaje de tiempo total de ejecución (usando `time.Now()` y `time.Since()`) al finalizar la función.
   - Manejar el error con el patrón canónico `if err != nil`.
   - Si todo es exitoso, imprimir las métricas con formateo claro usando `fmt.Printf`.
*/

func main() {
	// Escribe tu solución aquí
	start := time.Now()
	defer func() {
		fmt.Printf("Termino de ejecución: %v\n", time.Since(start))
	}()

	argumentos := os.Args[1:]

	if len(argumentos) == 0 {
		fmt.Println("Error: Debe ingresar números.")
		os.Exit(1)
	}

	numeros, err := parseInput(argumentos)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	min, max, avg, err := calculateStats(numeros)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al calcular %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Resultados:\n")
	fmt.Printf("- Valor mínimo: %.2f\n", min)
	fmt.Printf("- Valor máximo: %.2f\n", max)
	fmt.Printf("- promedio: %.2f\n", avg)
}

func calculateStats(numbers []float64) (min float64, max float64, avg float64, err error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("El slice está vacío")
	}

	min = numbers[0]
	max = numbers[0]
	suma := 0.0

	for _, number := range numbers {
		suma += number

		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}
	avg = suma / float64(len(numbers))
	return min, max, avg, nil
}

func parseInput(args []string) ([]float64, error) {
	numeros := make([]float64, len(args))

	for i, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, fmt.Errorf("error en el argumento '%s' (posición %d): %w", arg, i+1, err)
		}
		numeros[i] = num
	}
	return numeros, nil
}

//  go run main.go 10 20.5 30
