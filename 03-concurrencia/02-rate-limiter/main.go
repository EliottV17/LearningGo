package main

import (
	"context"
	"fmt"
	"sync"
	"time"

)

/*
EJERCICIO 3.2: Rate Limiter Token Bucket y Multiplexación con Context

OBJETIVO:
Manejar timeouts, cancelaciones estructuradas con `context.Context`,
multiplexación no bloqueante con `select` y control de flujo basado en tiempo (`time.Ticker`).

REQUERIMIENTOS:
1. Implementar un Rate Limiter con patrón Token Bucket:
   - Crear un canal de tokens con buffer `tokenBucket := make(chan time.Time, burstLimit)`.
   - Llenar inicialmente el canal con el límite de ráfaga (ej: 3 tokens).
   - Iniciar una goroutine que con un `time.NewTicker(interval)` añada tokens periódicamente al bucket sin bloquearse si está lleno (usando `select` con `default`).

2. Definir una función consumidora `executeRequest(ctx context.Context, reqID int, tokenBucket <-chan time.Time) error`:
   - Usar una sentencia `select` para esperar una de dos condiciones:
     a) `case <-tokenBucket`: se obtuvo un token disponible; simular la ejecución de la petición y retornar `nil`.
     b) `case <-ctx.Done()`: el contexto expiró o fue cancelado antes de obtener el token; retornar inmediatamente `ctx.Err()`.

3. En `main()`:
   - Crear un contexto raíz con timeout global usando `context.WithTimeout(context.Background(), totalTimeout)`.
   - Simular una ráfaga de 10 peticiones concurrentes en goroutines.
   - Usar `sync.WaitGroup` para esperar a que todas las peticiones terminen o aborten.
   - Contar cuántas peticiones fueron exitosas y cuántas fallaron por deadline/timeout.
   - Comprobar que ninguna goroutine queda bloqueada en memoria (zero goroutine leaks).
*/

func executeRequest(ctx context.Context, reqID int, tokenBucket <-chan time.Time) error {
	select {
	case <-tokenBucket:
		fmt.Printf("[Petición %d] Token obtenido exitosamente. Ejecutando...\n", reqID)
		time.Sleep(time.Millisecond * 50)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	// Escribe tu solución aquí
	burstLimit := 3
	tokenBucket := make(chan time.Time, burstLimit)

	for range burstLimit {
		tokenBucket <- time.Now()
	}

	interval := 500 * time.Millisecond
	ticker := time.NewTicker(interval)

	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				select {
				case tokenBucket <- t:
				default:
				}
			}
		}
	}()

	totalTimeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), totalTimeout)
	defer cancel()

	var wg sync.WaitGroup
	var mu sync.Mutex
	exitosas := 0
	fallidas := 0

	peticiones := 10

	for i := range peticiones{
		reqID := i + 1

		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			err := executeRequest(ctx, id, tokenBucket)

			if err != nil {
				fmt.Printf("Petición %d abortada: %v\n", id, err)
				mu.Lock()
				fallidas++
				mu.Unlock()
			} else {
				fmt.Printf("Petición %d completada con éxito.\n", id)
				mu.Lock()
				exitosas++
				mu.Unlock()
			}
		}(reqID)
	}

	wg.Wait()
	close(done)
	fmt.Println("\n--- REPORTE FINAL ---")
	fmt.Printf("Peticiones exitosas: %d\n", exitosas)
	fmt.Printf("Peticiones fallidas: %d\n", fallidas)
}
