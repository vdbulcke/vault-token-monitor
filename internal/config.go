package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
)

type Config struct {
	// Vault Config
	Address string `yaml:"vault_address" validate:"required"`
	Token   string `yaml:"vault_token" validate:"required"`

	// TLS
	SkipTLSValidation bool   `yaml:"skip_tls_validation"`
	CACertPEMPath     string `yaml:"vault_ca_pem_file" `

	// Prometheus metrics port
	PrometheusListeningPort int `yaml:"prometheus_listening_port" validate:"required"`

	// Scheduler
	SchedulerPeriodDuration string `yaml:"scheduler_period_duration" validate:"required"`

	// Vault Accessor Token List
	VaultAccessorMonitors []*VaultAccessorMonitor `yaml:"vault_accessor_token_list" validate:"required"`
}

type VaultAccessorMonitor struct {
	Accessor                   string `yaml:"token_accessor" validate:"required"`
	AutoRenewThresholdDuration string `yaml:"auto_renew_threshold_duration" `
}

// ValidateConfig validate config
func ValidateConfig(config *Config) bool {

	validate := validator.New()
	errs := validate.Struct(config)

	// validate internal token structure
	tokenError := false
	for _, a := range config.VaultAccessorMonitors {

		aErrs := validate.Struct(a)
		if aErrs != nil {
			tokenError = true
			for _, e := range aErrs.(validator.ValidationErrors) {
				fmt.Println(e)
			}
		}

		if a.AutoRenewThresholdDuration != "" {
			//nolint
			_, err := time.ParseDuration(a.AutoRenewThresholdDuration)
			if err != nil {
				fmt.Println(err.Error())
				tokenError = true
			}
		}

	}

	if errs == nil && !tokenError {
		return true
	}

	if errs != nil {
		for _, e := range errs.(validator.ValidationErrors) {
			fmt.Println(e)
		}
	}

	return false

}

// ParseConfig Parse config file
func ParseConfig(configFile string) (*Config, error) {

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	config := Config{}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	// return Parse config struct
	return &config, nil

}
