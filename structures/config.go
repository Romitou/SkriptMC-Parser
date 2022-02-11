package structures

import "strconv"

type Config struct {
	Version string
	HTTP    struct {
		Listen      string
		Port        int
		RateLimiter string
	}
	RateLimiter struct {
		Limit  string
		Prefix string
	}
	Redis struct {
		Address  string // takes env variable
		Username string // takes env variable
		Password string // takes env variable
		DB       int
	}
}

func (config Config) HTTPAddress() string {
	return config.HTTP.Listen + ":" + strconv.Itoa(config.HTTP.Port)
}
