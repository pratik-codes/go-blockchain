package utils

import "sync"

// helper function to execute the logic with lock
func WithLock(mu *sync.Mutex, action func()) {
	mu.Lock()         // Lock before executing the action
	defer mu.Unlock() // Ensure the mutex is unlocked after action is executed
	action()          // Execute the passed logic
}
