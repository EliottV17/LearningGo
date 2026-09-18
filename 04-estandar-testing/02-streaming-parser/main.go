package main

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
   - Si la entrada es un stream de objetos JSON delimitados por saltos de línea (NDJSON / JSON Lines) o un array JSON, procesar elemento por elemento sin cargar todo el contenido a memoria.
   - Contar cuántos registros tienen `Level == "ERROR"`.
   - Retornar el conteo total.

3. En `main()`:
   - Probar la función pasando primero un `strings.NewReader(...)` con datos de prueba hardcodeados (demostrando cómo cualquier tipo que implemente `Read` funciona).
   - Crear un archivo temporal con miles de líneas y procesarlo pasando el `*os.File` devuelto por `os.Open`.
   - Comprobar que el consumo de memoria se mantiene constante sin importar el tamaño del archivo (O(1) memory complexity).
*/

func main() {
	// Escribe tu solución aquí
}
