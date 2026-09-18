package main

import (
	"fmt"
)

/*
EJERCICIO 1.2: Disección de Memoria de Slices y Backing Arrays

OBJETIVO:
Entender la diferencia fundamental entre el header de un slice en Go (puntero, len, cap)
frente a los arrays dinámicos de JavaScript/V8. Demostrar cuándo ocurre una mutación compartida
y cuándo ocurre una reasignación con copia por exceder la capacidad.

REQUERIMIENTOS:
1. Crear una función `inspectSlice(label string, s []int)` que imprima:
   - La etiqueta (label).
   - Longitud (`len(s)`).
   - Capacidad (`cap(s)`).
   - Puntero al primer elemento del backing array (usando `fmt.Printf("%p\n", s)` o `&s[0]`).
   - Los elementos actuales.

2. En `main()`:
   - Crear un slice inicial `original` usando `make([]int, 3, 5)` (longitud 3, capacidad 5) e inicializarlo con valores [10, 20, 30].
   - Inspeccionar `original`.

   - Crear un subslices `sub1 := original[1:3]` (elementos [20, 30]).
   - Modificar `sub1[0] = 999`.
   - Inspeccionar ambos (`original` y `sub1`) y verificar cómo la mutación afectó al backing array compartido.

   - Hacer `append` a `sub1` con nuevos elementos hasta exceder su capacidad original.
   - Observar el momento exacto en que la dirección de memoria de `sub1` cambia (crecimiento geométrico del backing array).
   - Modificar nuevamente un elemento de `sub1` y comprobar que `original` ya NO cambia porque `sub1` ahora apunta a un nuevo backing array en el heap.

3. Documentar en comentarios en tu código la respuesta a:
   ¿Por qué en Go `append` siempre devuelve un nuevo slice header y qué pasaría si no reasignamos el resultado (`s = append(s, val)`)?
		Un slice en Go es un struct de 24 bytes(header) con tres campos: { ptr *T, len int, cap int}.
		Las funciones en Go reciben argumentos estrictamente por valor (copia).
		Cuando llamamos a `append(s, val)`:
		1. Si aún hay capacidad, `append` actualiza el elemento en el backing array y devuelve un nuevo
		header con `len + 1`. Si no reasignamos (`s = append(s, val)`), la variable local `s` conserva
		su `len` viejo y jamás se entera del nuevo elemento.
		2. Si se excede la capacidad (`len + nuevos > cap`), Go aloca un nuevo array en el heap con mayor
		capacidad, copia los elementos viejos y devuelve un nuevo header con un `ptr` que apunta a la nueva
		dirección de memoria. Si no reasignamos, seguimos apuntado al array viejo desactualizado.
*/

func main() {
	// // Escribe tu solución aquí
	// original := make([]int, 3, 5)
	// original[0] = 10
	// original[1] = 20
	// original[2] = 30
	// inspectSlice("Slice Original inicial", original)
	//
	// sub1 := original[1:3]
	// sub1[0] = 999
	// inspectSlice("Slice original tras mutar sub1", original)
	// inspectSlice("Slice sub1 tras mutar", sub1)
	//
	// sub2 := append(sub1, 40, 45, 50, 60)
	// inspectSlice("Slice sub2 (nuevo backing array)", sub2)
	// sub2[1] = 777
	//
	// inspectSlice("Slice sub1 tras sub2[1]", sub2)
	// inspectSlice("Slice Original(debe seguir intacto)", original)
	slice := []int{20, 30}
	nuevoElemento := 10
	slice = append([]int{nuevoElemento}, slice...)
	fmt.Println(slice)
}

func inspectSlice(label string, s []int) {
	fmt.Printf("[%s] -> dirección: %p | len %d | cap %d | datos: %v\n", label, s, len(s), cap(s), s)
}
