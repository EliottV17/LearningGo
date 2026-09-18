package main

/*
EJERCICIO 2.2: Cadena de Errores Tipados y Wrapping Idiomático

OBJETIVO:
Dominar el modelo de errores de Go como valores de primera clase,
implementando errores centinela (sentinel errors), structs de error personalizados,
wrapping con `%w`, e inspección con `errors.Is` y `errors.As`.

REQUERIMIENTOS:
1. Declarar dos errores centinela en package scope:
   - `var ErrNotFound = errors.New("resource not found")`
   - `var ErrUnauthorized = errors.New("unauthorized action")`

2. Crear un struct de error personalizado `ValidationError`:
   - Campos: `Field string`, `Reason string`.
   - Implementar el método `Error() string` para cumplir la interfaz nativa `error`.

3. Implementar un struct de error de contexto `DatabaseError`:
   - Campos: `Op string` (operación: "query", "connect", etc.), `Table string`, `Err error` (error interno subyacente).
   - Implementar `Error() string`.
   - Implementar `Unwrap() error` que retorne `d.Err` para habilitar el unwrap de la librería estándar.

4. Crear una función de simulación `processUser(userID string, role string) error`:
   - Si `userID == ""`, retornar un `ValidationError` con `Field: "userID"`, `Reason: "cannot be empty"`.
   - Si `role != "admin"`, envolver `ErrUnauthorized` con información adicional usando `fmt.Errorf("access denied for user %s: %w", userID, ErrUnauthorized)`.
   - Si `userID == "404"`, simular una falla de base de datos retornando un `DatabaseError` cuyo campo `Err` sea `ErrNotFound`.
   - Si todo es correcto, retornar `nil`.

5. En `main()`:
   - Invocar `processUser` con los distintos escenarios de prueba.
   - Para cada error recibido:
     a) Usar `errors.Is` para verificar si la causa raíz coincide con `ErrUnauthorized` o `ErrNotFound`.
     b) Usar `errors.As` para extraer el struct `ValidationError` e imprimir los campos específicos `Field` y `Reason`.
*/

func main() {
	// Escribe tu solución aquí
}
