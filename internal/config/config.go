package config

type Config struct {
	URL      string
	Users    int
	Rate     int // Requests per second limit
	Duration int
}

func NewConfig(url string, users, rate, duration int) *Config {
	return &Config{
		URL:      url,
		Users:    users,
		Rate:     rate,
		Duration: duration,
	}
}
