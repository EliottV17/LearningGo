package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

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

type User struct {
	ID    string
	Name  string
	Email string
}

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*User, error)
	Save(ctx context.Context, u *User) error
}

type InMemoryUserRepo struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]*User),
	}
}

func (repo *InMemoryUserRepo) Save(ctx context.Context, u *User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.users[u.ID] = u
	return nil
}

func (repo *InMemoryUserRepo) FindByID(ctx context.Context, id string) (*User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	user, exists := repo.users[id]
	if !exists {
		return nil, errors.New("usuario no encontrado")
	} else {
		return user, nil
	}
}

type UserService struct{ repo UserRepository }

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, name, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("el nombre no puede estar vacío")
	}

	if email == "" {
		return nil, errors.New("el email no puede estar vacío")
	}

	nuevoUsuario := &User{
		ID:    fmt.Sprintf("user-%d", time.Now().Unix()),
		Name:  name,
		Email: email,
	}

	err := s.repo.Save(ctx, nuevoUsuario)
	if err != nil {
		return nil, errors.New("error al guardar usuario")
	}
	return nuevoUsuario, nil
}

func main() {
	// Escribe tu solución aquí
	repo := NewInMemoryUserRepo()

	service := NewUserService(repo)

	ctx := context.Background()

	fmt.Println("Registrando usuario...")
	usuarioCreado, err := service.RegisterUser(ctx, "Gopher", "gopher@golang.org")
	if err != nil {
		fmt.Println("Error:", err)
		return 
	}
	fmt.Printf("Usuario %s registrado con ID: %s!\n", usuarioCreado.Name, usuarioCreado.ID)

	usuarioEncontrado, err := repo.FindByID(ctx, usuarioCreado.ID)
	if err != nil {
		fmt.Println("Error buscando:", err)
	} else {
		fmt.Printf("Usuario encontrado en BD: %s (%s)\n", usuarioEncontrado.Name, usuarioEncontrado.Email)
}
}
