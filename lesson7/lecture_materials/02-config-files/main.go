package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type LogLevel string

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host     string
	Port     int
	LogLevel LogLevel `yaml:"log_level" json:"log_level"`
}

func main() {
	configFile := flag.String("config", "", "application configuration file")

	flag.Parse()

	if configFile == nil || *configFile == "" {
		log.Fatal("config file not specified")
	}

	bytes, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("read config file error: %s", err)
	}

	cfg := Config{}

	if err := json.Unmarshal(bytes, &cfg); err != nil {
		log.Fatalf("broken config file: %s", err)
	}

	fmt.Printf("host: %s, port: %d, log level: %s\n", cfg.Server.Host, cfg.Server.Port, cfg.Server.LogLevel)
}
