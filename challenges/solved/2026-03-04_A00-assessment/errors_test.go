package main

import (
	"testing"
)

// - If `port` is not a valid integer → error
// - If `port` < 1 or > 65535 → error
// - Missing `host` → error
// - Unknown keys → ignore (don't error)
// - Use `fmt.Errorf` with `%w` for wrapping where appropriate

func TestParseConfig(t *testing.T) {

	t.Run("Parsing config object from data", func(t *testing.T) {
		data := `host=localhost
		port=8080
		debug=true`
		want := Config{Host: "localhost", Port: 8080, Debug: true}
		got, err := ParseConfig(data)

		if err != nil {
			t.Errorf("error parsing data: %s ", err)
		}
		if want != got {
			t.Errorf("objects do not match want %v got %v ", want, got)
		}
	})

}

func TestParseConfigValidation(t *testing.T) {

	test_cases := []struct {
		name        string
		input       string
		should_pass bool
		want        Config
	}{
		{name: "port is not valid integer", input: `
		host=localhost
		port=0
		debug=true
		`, should_pass: false},
		{name: "port is not valid integer",
			input: `host=localhost
			port=65536
			debug=true
		`, should_pass: false},
		{name: "missing host",
			input: `host=
					port=65534
					debug=true`,
			should_pass: false},
		{name: "debug is not valid boolean",
			input: `host=localhost
			port=65534
			debug=ok`,
			should_pass: false},
		{name: "debug is not valid boolean",
			input: `host=localhost
			port=8080
			debug=true
			nunknown_field=somevalue`, should_pass: true, want: Config{Host: "localhost", Port: 8080, Debug: true}},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := ParseConfig(tc.input)
			if tc.should_pass {
				if cfg != tc.want {
					t.Errorf("Test case %s failed, expect %v got %v", tc.name, tc.want, cfg)
				}
			} else {
				if err == nil {
					t.Errorf("Test case %s failed, expect should pass %v got %v", tc.name, tc.should_pass, err == nil)
				}
			}
		})
	}

}

// - Return a **custom error type** `ConfigError` that includes the field name and the reason
// - If `port` is not a valid integer → error
// - If `port` < 1 or > 65535 → error
// - Missing `host` → error
// - Unknown keys → ignore (don't error)
// - Use `fmt.Errorf` with `%w` for wrapping where appropriate
