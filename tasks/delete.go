package tasks

import (
	"net/http"
	"strings"
	"testWorkmate/storage"
)

// Обработка удаления задачи
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	if _, exists := storage.Tasks[id]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	delete(storage.Tasks, id)

	w.WriteHeader(http.StatusNoContent)
}
