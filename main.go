package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testWorkmate/tasks"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/tasks", tasks.CreateTaskHandler)
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tasks.GetTaskHandler(w, r)
		} else if r.Method == http.MethodDelete {
			tasks.DeleteTaskHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
