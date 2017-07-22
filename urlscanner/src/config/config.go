package config

import (
	"flag"
	"os"
	"time"
)

// UServer server configuration values set from env or flags
type UServer struct {
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	Version        string
}

// UCache server cache configuration values set from env or flags
type UCache struct {
	Addr     string
	Database int
	PassWord string
}

// Config server full configuration objects set from env or flags
type Config struct {
	UServer
	UCache
}

var config *Config

func init() {
	userver := UServer{
		Port:           ":9090",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Version:        "1",
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":9090"
	}
	version := os.Getenv("VERSION")
	if version == "" {
		version = "1"
	}

	flag.StringVar(&(userver.Port), "port", port, "Server port to run")
	flag.StringVar(&(userver.Version), "version", version, "Server running version")
	flag.Parse()
	ucache := UCache{
		Addr:     os.Getenv("CACHE_URL"), // "localhost:6379"
		PassWord: "",
		Database: 0,
	}
	config = &Config{
		UServer: userver,
		UCache:  ucache,
	}
}

// Cache get server cache configuration
func Cache() UCache {
	return config.UCache
}

// Get server  configuration
func Get() UServer {
	return config.UServer
}
