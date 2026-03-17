package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Host  string
	Port  int
	Debug bool
}

type ConfigError struct {
	Field string
	Err   error
}

func (cer *ConfigError) Error() string {
	return fmt.Sprintf("Error field %s, %v", cer.Field, cer.Err)
}

func (cer *ConfigError) Unwrap() error {
	return cer.Err
}

func ParseConfig(data string) (Config, error) {

	reader := strings.NewReader(data)
	scanner := bufio.NewScanner(reader)

	var cfg Config = Config{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		k, v, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		key := strings.ToLower(k)
		value := strings.ToLower(v)

		switch key {
		case "host":
			if value == "" {
				return cfg, &ConfigError{Field: key, Err: errors.New("missing value")}
			}
			cfg.Host = value
		case "port":
			if value == "" {
				return cfg, &ConfigError{Field: key, Err: errors.New("missing value")}
			}
			int_val, err := strconv.Atoi(value)
			if err != nil {
				return cfg, &ConfigError{Field: key, Err: fmt.Errorf("Invalid integer %w, unable to convert", err)}
			}
			if int_val < 1 || int_val > 65535 {
				return cfg, &ConfigError{Field: key, Err: fmt.Errorf("Invalid port range expect 1 ... 65535, got: %d", int_val)}
			}
			cfg.Port = int_val
		case "debug":
			bool_val, err := strconv.ParseBool(value)
			if err != nil {
				return cfg, &ConfigError{Field: key, Err: fmt.Errorf("Invalid type for boolean conversion %w", err)}
			}
			cfg.Debug = bool_val
		default:
			continue
		}
	}
	if scanner.Err() != nil {
		return cfg, errors.New("Error reading file")
	}
	return cfg, nil
}
