package main

import "testing"

/*
EJERCICIO 4.1 (TESTS): Table-Driven Tests y net/http/httptest

OBJETIVO:
Escribir tests idiomáticos en Go sin frameworks externos (sin Jest, sin Mocha, sin assert libraries).
Aplicar Table-Driven Tests y el paquete `net/http/httptest`.

REQUERIMIENTOS:
1. Implementar `TestGetItemHandler(t *testing.T)` usando una tabla de casos (Table-Driven):
   - Definir un struct anónimo para los casos de prueba:
     `tests := []struct {
         name           string
         itemID         string
         prepopulate    bool
         expectedStatus int
         expectedName   string
     }{ ... }`
   - Iterar los casos usando `for _, tc := range tests` y ejecutar `t.Run(tc.name, func(t *testing.T) { ... })`.

2. En cada subtest:
   - Crear una petición de prueba usando `httptest.NewRequest("GET", "/items/"+tc.itemID, nil)`.
   - Crear un recorder de respuesta usando `httptest.NewRecorder()`.
   - Ejecutar el handler/mux.
   - Comparar el código de estado: si `res.Code != tc.expectedStatus`, usar `t.Fatalf("expected status %d, got %d", tc.expectedStatus, res.Code)`.
   - Si se esperaba 200 OK, decodificar el body y verificar que los campos coincidan con los esperados.

3. Ejecutar los tests con: `go test -v -cover .`
*/

func TestGetItemHandler(t *testing.T) {
	// Escribe tus table-driven tests aquí
}
