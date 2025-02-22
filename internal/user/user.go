package user

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/8ideaz/hornet/internal/task"
)

// User simulates a single user running tasks
type User struct {
	Tasks  []*task.Task
	Client *http.Client
}

func NewUser(url string) *User {
	return &User{
		Tasks:  []*task.Task{task.SimpleTask(url)}, // Default task
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

// Run executes tasks randomly until stopped
func (u *User) Run(reportChan chan<- *TaskResult, stopChan <-chan struct{}) {
	for {
		select {
		case <-stopChan:
			return
		default:
			task := u.Tasks[0] // For now, just one task; we’ll add weighted selection later
			duration, err := task.Action(u.Client)
			result := &TaskResult{
				Name:     task.Name,
				Duration: duration,
				Error:    err != nil,
			}
			select {
			case reportChan <- result:
			case <-stopChan:
				return
			}
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) // Simulate think time
		}
	}
}

type TaskResult struct {
	Name     string
	Duration time.Duration
	Error    bool
}
