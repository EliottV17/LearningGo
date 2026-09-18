package main

/*
EJERCICIO 5.2: Graceful Shutdown en Servidores de Producción

OBJETIVO:
Manejar el ciclo de vida de una aplicación en contenedores (Docker/Kubernetes).
Interceptar señales del sistema operativo (SIGINT, SIGTERM) y cerrar conexiones
activas ordenadamente sin abortar peticiones en vuelo.

REQUERIMIENTOS:
1. Configurar un servidor `http.Server` en `:8080`:
   - Endpoint `/slow`: simula una operación larga durmiendo por 5 segundos (`time.Sleep(5 * time.Second)`) antes de responder "Operación finalizada".
   - Endpoint `/healthz`: responde inmediatamente con status 200 "OK".

2. Configurar la captura de señales del SO:
   - Usar `signal.Notify` con un canal `os.Signal` escuchando `syscall.SIGINT` y `syscall.SIGTERM`,
     o la función moderna `signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)`.

3. Orquestación del ciclo de vida en `main()`:
   - Arrancar el servidor en una goroutine secundaria con `server.ListenAndServe()`.
   - Capturar si retorna un error distinto a `http.ErrServerClosed`.
   - En la goroutine principal, bloquearse esperando la señal de interrupción.
   - Al recibir la señal, registrar el evento ("Iniciando apagado ordenado...") y crear un contexto con timeout de 10 segundos:
     `shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)`
     `defer cancel()`
   - Invocar `server.Shutdown(shutdownCtx)`.
   - Manejar el cierre de recursos adicionales (bases de datos simuladas, pools de workers, etc.).
   - Registrar la finalización exitosa del proceso.

4. Prueba manual:
   - Iniciar el programa.
   - Enviar una petición a `/slow` en una terminal (`curl http://localhost:8080/slow`).
   - Inmediatamente presionar `Ctrl+C` en la terminal del servidor.
   - Verificar que el servidor NO se muere de golpe: espera a que `curl` termine de recibir su respuesta y luego se apaga limpiamente.
*/

func main() {
	// Escribe tu solución aquí
}
