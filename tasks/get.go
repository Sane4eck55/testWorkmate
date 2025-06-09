package tasks

import (
	"encoding/json"
	"net/http"
	"strings"
	"testWorkmate/storage"
)

// Обработка получения статуса задачи
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "Missing task ID", http.StatusBadRequest)
		return
	}

	storage.TasksMu.Lock()
	defer storage.TasksMu.Unlock()
	task, exists := storage.Tasks[id]

	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
