package storage

import (
	"sync"
	"testWorkmate/datastruct"
)

// Хранилище задач
var (
	Tasks   = make(map[string]*datastruct.Task)
	TasksMu sync.Mutex
)
