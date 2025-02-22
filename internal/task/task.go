package task

import (
	"net/http"
	"time"
)

// Task represents a single user action (e.g., an HTTP request)
type Task struct {
	Name   string
	Weight int // For task selection probability (like Locust)
	Action func(*http.Client) (time.Duration, error)
}

// SimpleTask creates a basic GET request task
func SimpleTask(url string) *Task {
	return &Task{
		Name:   "GET " + url,
		Weight: 1,
		Action: func(client *http.Client) (time.Duration, error) {
			start := time.Now()
			req, _ := http.NewRequest("GET", url, nil)
			resp, err := client.Do(req)
			if err != nil {
				return 0, err
			}
			resp.Body.Close()
			return time.Since(start), nil
		},
	}
}
