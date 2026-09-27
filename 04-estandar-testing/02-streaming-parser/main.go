package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
EJERCICIO 4.2: Streaming I/O y Decodificación JSON con io.Reader

OBJETIVO:
Abandonar la costumbre de JS de cargar archivos enteros o payloads completos en memoria (como `fs.readFileSync` o `await req.json()`).
Aprender la potencia de las interfaces fundamentales de Go: `io.Reader`, `io.Writer` y `json.Decoder`.

REQUERIMIENTOS:
1. Modelar un evento de log:
   - `LogEntry`: `Timestamp string `json:"timestamp"``, `Level string `json:"level"``, `Message string `json:"message"``

2. Implementar una función `countErrors(r io.Reader) (int, error)`:
   - Recibir una abstracción `io.Reader` (no una ruta de archivo `string`, ni un `[]byte`).
   - Crear un `json.NewDecoder(r)`.
   - Si la entrada es un stream de objetos JSON delimitados por saltos de línea (NDJSON / JSON Lines) o un array JSON
		, procesar elemento por elemento sin cargar todo el contenido a memoria.
   - Contar cuántos registros tienen `Level == "ERROR"`.
   - Retornar el conteo total.

3. En `main()`:
   - Probar la función pasando primero un `strings.NewReader(...)` con datos de prueba hardcodeados (demostrando cómo cualquier tipo que implemente `Read` funciona).
   - Crear un archivo temporal con miles de líneas y procesarlo pasando el `*os.File` devuelto por `os.Open`.
   - Comprobar que el consumo de memoria se mantiene constante sin importar el tamaño del archivo (O(1) memory complexity).
*/

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

func countErrors(r io.Reader) (int, error) {
	contadorErrores := 0

	decoder := json.NewDecoder(r)

	for {
		var entry LogEntry

		err := decoder.Decode(&entry)

		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		if entry.Level == "ERROR" {
			contadorErrores++
		}
	}
	return contadorErrores, nil
}

func main() {
	// Escribe tu solución aquí

	datosDePrueba := `
    {"timestamp":"2024-01-01T10:00:00Z", "level":"INFO", "message":"Iniciando sistema"}
    {"timestamp":"2024-01-01T10:01:00Z", "level":"ERROR", "message":"Fallo de conexión"}
    {"timestamp":"2024-01-01T10:02:00Z", "level":"WARN", "message":"Reintentando"}
    {"timestamp":"2024-01-01T10:03:00Z", "level":"ERROR", "message":"Timeout de base de datos"}
  `

	lectorMemoria := strings.NewReader(datosDePrueba)

	errores, err := countErrors(lectorMemoria)
	if err != nil {
		fmt.Println("Error procesando:", err)
	} else {
		fmt.Println("Errores encontrados (Memoria):", errores)
	}

	archivoPrueba, _ := os.Create("logs_gigantes.json")

	for i := range 10000 {
		if i%2 == 0 {
			archivoPrueba.WriteString(`{"timestamp":"2024-01-01T10:00:00Z", "level":"INFO", "message":"Todo OK"}` + "\n")
		} else {
			archivoPrueba.WriteString(`{"timestamp":"2024-01-01T10:00:00Z", "level":"ERROR", "message":"Fallo grave"}` + "\n")
		}
	}
	archivoPrueba.Close()

	file, err := os.Open("logs_gigantes.json")
	if err != nil {
		fmt.Println("Error abriendo el archivo:", err)
		return
	}

	defer file.Close()

	erroresArchivo, err := countErrors(file)
	if err != nil {
		fmt.Println("Error procesando archivo:", err)
	} else {
		fmt.Println("Errores encontrados (Archivos físico):", erroresArchivo)
	}
}
