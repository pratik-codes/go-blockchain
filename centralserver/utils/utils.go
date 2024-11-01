package utils

import (
	"fmt"
	"sync"
)

func WithLock(mu *sync.Mutex, action func()) {
	mu.Lock()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in action:", r)
		}
		mu.Unlock()
	}()
	action()
}
