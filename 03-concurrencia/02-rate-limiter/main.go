package main

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

func main() {
	// Escribe tu solución aquí
}
