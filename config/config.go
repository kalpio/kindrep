package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Host       `json:"host"`
	Credential `json:"credential"`
	Sender     EmailAddress   `json:"sender"`
	Receivers  []EmailAddress `json:"receivers"`
}

type Credential struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Host struct {
	HostName string `json:"hostName"`
	Port     uint16 `json:"port"`
}

type EmailAddress struct {
	Address string `json:"address"`
}

func Read(fileName string) (*Config, error) {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("kindrep: %w", err)
	}
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		return nil, fmt.Errorf("kindrep: %w", err)
	}

	return readData(f)
}

func readData(reader io.Reader) (*Config, error) {
	result := &Config{}
	if err := json.NewDecoder(reader).Decode(&result); err != nil {
		return nil, fmt.Errorf("kindrep: %w", err)
	}
	return result, nil
}
