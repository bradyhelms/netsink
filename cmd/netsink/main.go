package main

import (
	"fmt"
	"github.com/bradyhelms/netsink/pkg/config"
	"github.com/goccy/go-yaml"
	"log"
	"os"
	"path/filepath"
)

func readStateFile(filepath string) config.NetworkState {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var cfg config.NetworkState

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Failed to unmarshal yaml: %v", err)
	}

	fmt.Printf("Successfully loaded config: $+v\n", cfg)
	return cfg
}

func main() {
	path := filepath.Join("config", "desired_state.yaml")
	var desiredState config.NetworkState = readStateFile(path)
	path = filepath.Join("config", "drift_state.yaml")
	var currentState config.NetworkState = readStateFile(path)

	fmt.Printf("%v", desiredState)
	fmt.Printf("%v", currentState)
}
