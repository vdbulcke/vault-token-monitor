package internal

import "testing"

func TestConfig(t *testing.T) {

	configFile := "../example/config.yaml"

	config, err := ParseConfig(configFile)
	if err != nil {
		t.Fatal(err)
	}

	if !ValidateConfig(config) {
		t.Fatal(err)
	}

}
