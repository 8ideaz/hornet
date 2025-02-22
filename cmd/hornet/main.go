package main

import (
	"flag"
	"fmt"

	"github.com/8ideaz/hornet/internal/worker"

	"github.com/8ideaz/hornet/internal/config"
)

var (
	url      string
	users    int
	duration int
	rate     int
)

func main() {
	// Define command-line flags
	flag.StringVar(&url, "url", "http://localhost:8080", "Target URL for load testing")
	flag.IntVar(&users, "users", 10, "Number of concurrent users")
	flag.IntVar(&duration, "duration", 30, "Test duration in seconds")
	flag.IntVar(&rate, "rate", 100, "Requests per second limit (0 for no limit)")
	flag.Parse()

	// Load configuration
	cfg := config.NewConfig(url, users, rate, duration)

	// Run the load test
	fmt.Printf("Starting Hornet load test: %d users, %d seconds, targeting %s\n", cfg.Users, cfg.Duration, cfg.URL)
	results := worker.RunLoadTest(cfg)

	// Display results (metrics TBD)
	fmt.Println("Test completed. Results:")
	fmt.Printf("Total Requests: %d\n", results.TotalRequests)
}
