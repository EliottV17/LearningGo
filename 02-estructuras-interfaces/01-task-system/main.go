package main

import (
	"fmt"
)

/*
EJERCICIO 2.1: Sistema de Tareas Notificables (Composición e Interfaces Implícitas)

OBJETIVO:
Abandonar los paradigmas de clases y herencia de TypeScript/NestJS.
Aprender a modelar comportamiento mediante composición de structs e interfaces implícitas.

REQUERIMIENTOS:
1. Definir la interfaz `Notifier`:
   - Método: `Notify(message string) error`

2. Implementar dos structs concretos que cumplan la interfaz `Notifier`:
   - `EmailNotifier`: campos `Email string`, `Host string`.
   - `SlackNotifier`: campos `WebhookURL string`, `Channel string`.
   - Recuerda: NO existe palabra clave `implements`. El cumplimiento es 100% implícito si firmas y tipos coinciden.

3. Definir un struct `Task`:
   - Campos: `ID string`, `Title string`, `Completed bool`.
   - Métodos:
     - `Complete()`: Usa un **pointer receiver** (`(t *Task) Complete()`) para mutar `Completed` a `true`.
     - `String()`: Usa un **value receiver** (`(t Task) String() string`) para representar la tarea formateada. Cumple implícitamente `fmt.Stringer`.

4. Definir un struct `ProjectManager`:
   - Embeber un slice de `Task` y un campo de tipo `Notifier` (la interfaz).
   - Método `AddTask(title string) *Task`: crea una nueva tarea, la añade a la lista y retorna su puntero.
   - Método `FinishTask(id string) error`: busca la tarea, la completa llamando a su método y, si todo va bien,
		envía una notificación usando el `Notifier` inyectado.

5. En `main()`:
   - Instanciar `ProjectManager` primero con `EmailNotifier` y completar una tarea.
   - Instanciar otro `ProjectManager` con `SlackNotifier` y completar otra tarea.
   - Probar qué pasa si intentas llamar a `Complete()` sobre una copia por valor en lugar de un puntero.
*/

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct {
	Email string
	Host  string
}

func (e EmailNotifier) Notify(message string) error {
	fmt.Printf("[Email a %s vía %s]: %s\n", e.Email, e.Host, message)
	return nil
}

type SlackNotifier struct {
	WebhookURL string
	Channel    string
}

func (s SlackNotifier) Notify(message string) error {
	fmt.Printf("[Slack al canal %s vía %s]: %s\n", s.Channel, s.WebhookURL, message)
	return nil
}

type Task struct {
	ID        string
	Title     string
	Completed bool
}

func (t *Task) Complete() {
	t.Completed = true
}

func (t Task) String() string {
	return fmt.Sprintf("La tarea '%s' con el título '%s' ya esta completada: %v", t.ID, t.Title, t.Completed)
}

type ProjectManager struct {
	sliceTask []*Task
	Notifier
}

func (pm *ProjectManager) AddTask(title string) *Task {
	nuevaTarea := &Task{
		ID: fmt.Sprintf("%d", len(pm.sliceTask)+1),
		Title: title,
	}

	pm.sliceTask = append(pm.sliceTask, nuevaTarea)

	return nuevaTarea
}

func (pm *ProjectManager) FinishTask(id string) error {
	for _, tarea := range pm.sliceTask {
		if tarea.ID == id {
			tarea.Complete()

			mensaje := fmt.Sprintf("La tarea %s está lista", tarea.Title)
			pm.Notify(mensaje)
			return nil
		}
	}
	return fmt.Errorf("no se encontró la tarea con ID %s", id)
}

func main() {
	// Escribe tu solución aquí
	notificadorEmail := EmailNotifier{
		Email: "dev@empresa.com",
		Host: "smtp.gmail.com",
	}

	pm1 := ProjectManager{
		sliceTask: []*Task{},
		Notifier: notificadorEmail,
	}

	pm1.AddTask("Configurar el servidor")
	pm1.FinishTask("1")

	notificadorSlack := SlackNotifier{
		WebhookURL: "https://hooks.slack.com/...",
		Channel: "#backend-alerts",
	}

	pm2 := ProjectManager{
		sliceTask: []*Task{},
		Notifier: notificadorSlack,
	}

	pm2.AddTask("Migrar base de datos")
	pm2.FinishTask("1")

	fmt.Println("\n--- Prueba de Mutación ---")

	tareaOriginal := Task{ID: "99", Title: "Prueba de Valor", Completed: false}

	tareaCopia := tareaOriginal

	tareaCopia.Complete()

	fmt.Printf("Original completada? %v\n", tareaOriginal.Completed)
	fmt.Printf("Copia Completada? %v\n", tareaCopia.Completed)
}
