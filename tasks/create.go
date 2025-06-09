package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testWorkmate/datastruct"
	"testWorkmate/storage"
	"time"

	"golang.org/x/exp/rand"
)

const (
	minDuration = 60
)

// Обработка создания новой задачи
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := generateID()

	duration := rand.Intn(60)

	if duration == 0 {
		duration = minDuration
	}

	task := &datastruct.Task{
		ID:        id,
		Status:    datastruct.StatusPending,
		CreatedAt: time.Now(),
		Duration:  duration,
	}

	storage.TasksMu.Lock()
	defer storage.TasksMu.Unlock()
	storage.Tasks[id] = task

	go processTask(task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Генерация уникального ID для задачи
func generateID() string {
	return fmt.Sprintf("%d", rand.Uint64())
}

// Имитация выполнения задачи
func processTask(task *datastruct.Task) {
	task.Status = datastruct.StatusRunning
	task.StartTime = time.Now()

	time.Sleep(time.Duration(task.Duration) * time.Second)

	task.Status = datastruct.StatusCompleted
	task.EndTime = time.Now()
}
