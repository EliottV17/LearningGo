package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

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

type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type ItemStore struct {
	rwm  sync.RWMutex
	mapa map[string]Item
}

func (s *ItemStore) Save(item Item) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.mapa[item.ID] = item
}

func (s *ItemStore) Get(id string) (Item, bool) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	item, exists := s.mapa[id]
	if !exists {
		return Item{}, false
	}
	return item, true
}

func (s *ItemStore) List() []Item {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	items := make([]Item, 0, len(s.mapa))
	for _, item := range s.mapa {
		items = append(items, item)
	}
	return items
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()

		next.ServeHTTP(w, r)

		tiempoTranscurrido := time.Since(inicio)

		log.Printf("[%s] %s - duración: %s", r.Method, r.URL.Path, tiempoTranscurrido)
	})
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("!PÁNICO ATRAPADO!: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (s *ItemStore) handleListItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	items := s.List()

	err := json.NewEncoder(w).Encode(items)

	if err != nil {
		http.Error(w, "Error al procesar los datos", http.StatusInternalServerError)
	}
}

func (s *ItemStore) handleGetItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	item, found := s.Get(id)

	if !found {
		http.Error(w, "Elemento no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(item)
}

func (s *ItemStore) handleCreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item

	err := json.NewDecoder(r.Body).Decode(&newItem)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if newItem.ID == "" || newItem.Name == "" {
		http.Error(w, "Estos campos no pueden estar vacíos", http.StatusBadRequest)
		return
	}

	newItem.CreatedAt = time.Now()

	s.Save(newItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newItem)
}

func main() {
	// Escribe tu solución aquí
	store := &ItemStore{
		mapa: make(map[string]Item),
	}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /items", store.handleListItems)
	mux.HandleFunc("GET /items/{id}", store.handleGetItem)
	mux.HandleFunc("GET /items", store.handleCreateItem)

	handleConLogging := loggingMiddleware(mux)
	handleProtegido := recoverMiddleware(handleConLogging)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handleProtegido,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Servidor corriendo en http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
