package runner

import (
	"sync"
	"time"

	"github.com/8ideaz/hornet/internal/config"
	"github.com/8ideaz/hornet/internal/stats"
	"github.com/8ideaz/hornet/internal/user"
)

func Run(cfg *config.Config) {
	stats := stats.NewStats()
	reportChan := make(chan *user.TaskResult, 100)
	stopChan := make(chan struct{})
	var wg sync.WaitGroup

	// Spawn users
	for i := 0; i < cfg.Users; i++ {
		wg.Add(1)
		u := user.NewUser(cfg.URL)
		go func() {
			defer wg.Done()
			u.Run(reportChan, stopChan)
		}()
	}

	// Collect stats
	go func() {
		for result := range reportChan {
			stats.Record(result.Duration, result.Error)
		}
	}()

	// Rate limiting (simplified for now)
	go func() {
		time.Sleep(time.Duration(cfg.Duration) * time.Second)
		close(stopChan)
	}()

	// Wait for test to complete
	wg.Wait()
	close(reportChan)

	// Report results
	stats.Report()
}
