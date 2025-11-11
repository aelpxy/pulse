package config

import (
	"time"
)

type Config struct {
	Host string
	Port int

	AppKey    string
	AppSecret string

	MaxConnections            int
	MaxChannelsPerConnection  int
	MaxSubscriptionsPerSecond int

	ActivityTimeout  time.Duration
	WriteTimeout     time.Duration
	ReadTimeout      time.Duration
	HandshakeTimeout time.Duration
	PingInterval     time.Duration
	PongTimeout      time.Duration

	WriteBufferSize   int
	ReadBufferSize    int
	MessageBufferSize int

	EventsPerSecond int
	EventBurst      int

	AllowOrigins []string

	EnableMetrics bool
	EnablePprof   bool
	EnableRedis   bool

	// for later
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	RedisPoolSize int
}

func DefaultConfig() *Config {
	return &Config{
		Host: "0.0.0.0",
		Port: 8080,

		AppKey:    "",
		AppSecret: "",

		MaxConnections:            100000,
		MaxChannelsPerConnection:  100,
		MaxSubscriptionsPerSecond: 10,

		ActivityTimeout:  120 * time.Second,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      60 * time.Second,
		HandshakeTimeout: 5 * time.Second,
		PingInterval:     30 * time.Second,
		PongTimeout:      10 * time.Second,

		WriteBufferSize:   4096,
		ReadBufferSize:    4096,
		MessageBufferSize: 256,

		EventsPerSecond: 100,
		EventBurst:      200,

		AllowOrigins: []string{"*"},

		EnableMetrics: true,
		EnablePprof:   false,
		EnableRedis:   false,

		// temp values
		RedisAddr:     "localhost:6379",
		RedisPassword: "",
		RedisDB:       0,
		RedisPoolSize: 100,
	}
}
