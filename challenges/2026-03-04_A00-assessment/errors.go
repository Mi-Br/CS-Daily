package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Host  string
	Port  int64
	Debug bool
}

type ConfigError struct {
	field string
	err   error
}

func (cr ConfigError) Error() string {
	return fmt.Sprintf("%s is %s", cr.field, cr.err)
}

func ParseConfig(data string) (Config, error) {

	reader := strings.NewReader(data)
	scanner := bufio.NewScanner(reader)

	var cfg Config

	for scanner.Scan() {
		line := scanner.Text()
		k, v, found := strings.Cut(line, "=")
		if found {
			if k == "Port" {
				p_int, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return cfg, ConfigError{field: string(p_int), err: fmt.Errorf("Not valid integer")}
				}
				cfg.Port = p_int
			}
			if k == "Host" {
				if len(v) == 0 {
					return cfg, ConfigError{field: k, err: fmt.Errorf("missing")}
				}
				cfg.Host = v
			}
			if k == "Debug" {
				b, err := strconv.ParseBool(v)
				if err != nil {
					return cfg, ConfigError{field: k, err: fmt.Errorf("Not valid value for bool")}
				}
				cfg.Debug = b
			}
		}
		return cfg, nil
	}
}

// ### Exercise 4: Error Handling
// Write a function `ParseConfig(data string) (Config, error)` where `Config` is:

// ```go
// type Config struct {
//     Host    string
//     Port    int
//     Debug   bool
// }
// ```

// The input `data` is a simple format:
// ```
// host=localhost
// port=8080
// debug=true
// ```

// Rules:
// - Return a **custom error type** `ConfigError` that includes the field name and the reason
// - If `port` is not a valid integer → error
// - If `port` < 1 or > 65535 → error
// - Missing `host` → error
// - Unknown keys → ignore (don't error)
// - Use `fmt.Errorf` with `%w` for wrapping where appropriate

// ---
