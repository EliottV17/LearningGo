package main

/*
EJERCICIO 4.1: API REST Idiomática con net/http (Estilo Go 1.22+)

OBJETIVO:
Crear un servidor HTTP de producción usando exclusivamente la librería estándar (`net/http`),
implementando routing enriquecido (métodos y wildcards de Go 1.22+), middlewares componibles
y serialización JSON eficiente.

REQUERIMIENTOS:
1. Modelar la entidad `Item`:
   - `ID string `json:"id"``
   - `Name string `json:"name"``
   - `Price float64 `json:"price"``
   - `CreatedAt time.Time `json:"created_at"``

2. Implementar un `ItemStore` thread-safe:
   - Struct con `sync.RWMutex` y un mapa interno `map[string]Item`.
   - Métodos: `Get(id string) (Item, bool)`, `Save(item Item)`, `List() []Item`.

3. Implementar Middlewares estándar con la firma `func(http.Handler) http.Handler`:
   - `loggingMiddleware`: registra método, URL y duración de la petición.
   - `recoverMiddleware`: captura cualquier `panic` en el handler mediante `recover()`, registra el error y responde con HTTP 500 JSON.

4. Crear handlers HTTP utilizando `http.ServeMux` de Go 1.22+:
   - `"GET /items"`: lista todos los items en formato JSON.
   - `"GET /items/{id}"`: obtiene un item por ID usando `r.PathValue("id")`. Si no existe, responder 404.
   - `"POST /items"`: decodifica el body usando `json.NewDecoder(r.Body)` (evitar `io.ReadAll`), valida campos, guarda y responde 201 Created.

5. En `main()`:
   - Configurar `http.Server` con timeouts explícitos (`ReadTimeout`, `WriteTimeout`, `IdleTimeout`) para evitar ataques de resource exhaustion (slowloris).
   - Levantar el servidor en el puerto `:8080`.
*/

func main() {
	// Escribe tu solución aquí
}
