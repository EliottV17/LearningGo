package main

/*
EJERCICIO 5.1: Inyección de Dependencias Idiomática y Desacoplamiento

OBJETIVO:
Diseñar servicios desacoplados en Go sin recurrir a contenedores mágicos de IoC (como NestJS o InversifyJS).
En Go, la inyección de dependencias se realiza mediante interfaces y funciones constructoras (`New...`).
"Accept interfaces, return structs" (Acepta interfaces, retorna structs).

REQUERIMIENTOS:
1. Definir la interfaz de persistencia en el paquete consumidor:
   - `type UserRepository interface {
         FindByID(ctx context.Context, id string) (*User, error)
         Save(ctx context.Context, u *User) error
     }`
   - *Regla de oro de Go*: la interfaz pertenece a quien la consume, no a quien la implementa.

2. Implementar una implementación concreta en memoria:
   - `type InMemoryUserRepo struct { ... }`
   - Función constructora: `func NewInMemoryUserRepo() *InMemoryUserRepo`.

3. Implementar el servicio de negocio:
   - `type UserService struct { repo UserRepository }`
   - Función constructora: `func NewUserService(repo UserRepository) *UserService`.
   - Método `RegisterUser(ctx context.Context, name string, email string) (*User, error)`: valida datos, crea el usuario y lo persiste a través del repositorio.

4. En `main()`:
   - "Wirear" (ensamblar) manualmente las dependencias en la raíz de composición (`main`):
     `repo := NewInMemoryUserRepo()`
     `service := NewUserService(repo)`
   - Ejecutar operaciones de prueba con un `context.Background()`.
   - Notar la claridad absoluta de no tener decoradores `@Injectable()` ni resolución por reflexión en runtime.
*/

func main() {
	// Escribe tu solución aquí
}
