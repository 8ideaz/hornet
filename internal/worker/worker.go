package worker

import (
	"net/http"
	"sync"
	"time"

	"github.com/8ideaz/hornet/internal/config"
)

type Result struct {
	mu            sync.Mutex
	TotalRequests int
	// Add more metrics later (e.g., latency, errors)
}

func RunLoadTest(cfg *config.Config) Result {
	var wg sync.WaitGroup
	results := Result{}

	// Channel to collect results from workers (for simplicity, just counting requests now)
	done := make(chan struct{})

	// Start workers
	for i := 0; i < cfg.Users; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runWorker(cfg.URL, done, &results)
		}()
	}

	// Wait for duration, then signal workers to stop
	time.Sleep(time.Duration(cfg.Duration) * time.Second)
	close(done)

	// Wait for all workers to finish
	wg.Wait()
	return results
}

func runWorker(url string, done <-chan struct{}, results *Result) {
	client := &http.Client{Timeout: 10 * time.Second}
	for {
		select {
		case <-done:
			return
		default:
			req, _ := http.NewRequest("GET", url, nil)
			resp, err := client.Do(req)
			if err == nil {
				resp.Body.Close()
				// Safely increment total requests
				results.mu.Lock()
				results.TotalRequests++
				results.mu.Unlock()
			}
		}
	}
}
