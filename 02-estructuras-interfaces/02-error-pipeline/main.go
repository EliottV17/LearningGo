package main

import (
	"errors"
	"fmt"
)

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

var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized action")
)

type ValidationError struct {
	Field  string
	Reason string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("error de validación en el campo '%s': %s", ve.Field, ve.Reason)
}

type DatabaseError struct {
	Op    string
	Table string
	Err   error
}

func (d *DatabaseError) Error() string {
	return fmt.Sprintf("error en la base de datos en la operación %s y la tabla %s: %v", d.Op, d.Table, d.Err.Error())
}

func (d *DatabaseError) Unwrap() error {
	return d.Err
}

func processUser(userID string, role string) error {
	if userID == "" {
		return &ValidationError{
			Field:  "userID",
			Reason: "cannot be empty",
		}
	}

	if role != "admin" {
		return fmt.Errorf("acces denied for user %s: %w", userID, ErrUnauthorized)
	}

	if userID == "404" {
		return &DatabaseError{
			Op:    "query",
			Table: "users",
			Err:   ErrNotFound,
		}
	}

	return nil
}

func analyzeError(err error) {
	if err != nil {
		fmt.Println("Error recibido:", err)
	}

	if errors.Is(err, ErrNotFound) {
		fmt.Println("-> [Alerta] El usuario no fue encontrado en la base de datos.")
	}

	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("-> [Alerta] Intengo de acceso sin permisos.")
	}

	if valErr, ok := errors.AsType[*ValidationError](err); ok {
		fmt.Printf("-> [Validación Fallida] Campo afectado: %s, Motivo: %s\n", valErr.Field, valErr.Reason)
	}
	fmt.Println("---------------------------------------------")
}

func main() {
	// Escribe tu solución aquí
	analyzeError(processUser("","admin"))
	analyzeError(processUser("123","user"))
	analyzeError(processUser("404","admin"))
}
