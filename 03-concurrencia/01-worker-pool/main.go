package main

/*
EJERCICIO 3.1: Worker Pool Concurrente Robusto (CSP Pattern)

OBJETIVO:
Implementar el patrón canónico de Worker Pool en Go para procesar tareas concurrentes
con límite estricto de concurrencia, recolección de resultados, sincronización con `sync.WaitGroup`
y control de terminación limpia sin goroutine leaks.

REQUERIMIENTOS:
1. Modelar las estructuras de datos:
   - `Job`: `ID int`, `Data string`, `Cost time.Duration` (tiempo simulado de procesamiento).
   - `Result`: `JobID int`, `WorkerID int`, `Output string`, `Duration time.Duration`, `Err error`.

2. Implementar la función trabajadora `worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup)`:
   - Notar las direcciones de los channels en la firma: `jobs` es sólo lectura (`<-chan`), `results` es sólo escritura (`chan<-`).
   - Usar `defer wg.Done()` para garantizar que la finalización del worker se registre en el WaitGroup.
   - Procesar items de `jobs` mediante un bucle `for job := range jobs`.
   - Simular el cómputo con `time.Sleep(job.Cost)`.
   - Enviar el resultado estructurado al canal `results`.

3. En `main()`:
   - Definir una cantidad fija de workers (por ejemplo, 3 workers) y un lote de jobs (por ejemplo, 10 jobs).
   - Crear los canales: `jobs` (buffered con capacidad acorde) y `results` (buffered).
   - Levantar los workers en goroutines (`go worker(...)`).
   - Encolar los jobs en el canal `jobs` y cerrarlo inmediatamente (`close(jobs)`) para avisar a los workers que no habrá más trabajo.
   - Levantar una goroutine separada dedicada a esperar (`wg.Wait()`) y cerrar el canal `results` una vez que todos los workers terminen.
   - En el hilo principal de `main`, drenar e imprimir los resultados usando `for res := range results`.
   - Correr y verificar la solución con el race detector: `go run -race .`
*/

func main() {
	// Escribe tu solución aquí
}
