package stats

import (
	"fmt"
	"sync"
	"time"
)

type Stats struct {
	Requests  int
	Failures  int
	TotalTime time.Duration
	mu        sync.Mutex
}

func NewStats() *Stats {
	return &Stats{}
}

func (s *Stats) Record(duration time.Duration, isError bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Requests++
	if isError {
		s.Failures++
	}
	s.TotalTime += duration
}

func (s *Stats) Report() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Requests == 0 {
		fmt.Println("No requests recorded.")
		return
	}
	avgTime := s.TotalTime / time.Duration(s.Requests)
	fmt.Printf("Requests: %d, Failures: %d, Avg Response Time: %v\n", s.Requests, s.Failures, avgTime)
}
