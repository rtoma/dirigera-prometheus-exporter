package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type DirigeraExporterConfig struct {
	Url         string
	AccessToken string
	Interval    time.Duration
	BindAddress string
}

func NewConfig() (*DirigeraExporterConfig, error) {

	url := os.Getenv("DIRIGERA_URL")
	token := os.Getenv("DIRIGERA_ACCESS_TOKEN")

	if url == "" || token == "" {
		return nil, errors.New("DIRIGERA_URL and DIRIGERA_ACCESS_TOKEN must be set")
	}

	interval := os.Getenv("DIRIGERA_INTERVAL")
	var intervalSec int
	var err error

	if interval == "" {
		intervalSec = 30
	} else {
		intervalSec, err = strconv.Atoi(interval)
		if err != nil {
			return nil, fmt.Errorf("DIRIGERA_INTERVAL value is not a number: %s", interval)
		}
		if intervalSec < 5 || intervalSec > 120 {
			return nil, fmt.Errorf("DIRIGERA_INTERVAL value should be between 5 and 120 (seconds)")
		}
	}

	bindAddress := os.Getenv("DIRIGERA_BIND_ADDRESS")
	if bindAddress == "" {
		bindAddress = ":9090"
	}

	return &DirigeraExporterConfig{
		Url:         url,
		AccessToken: token,
		Interval:    time.Duration(intervalSec) * time.Second,
		BindAddress: bindAddress,
	}, nil
}
