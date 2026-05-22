package config

import (
	"fmt"
	"os"

	"github.com/go-core-fx/config"
)

type http struct {
	Address     string   `koanf:"address"`
	ProxyHeader string   `koanf:"proxy_header"`
	Proxies     []string `koanf:"proxies"`

	OpenAPI openAPIConfig `koanf:"openapi"`
}

type openAPIConfig struct {
	Enabled    bool   `koanf:"enabled"`
	PublicHost string `koanf:"public_host"`
	PublicPath string `koanf:"public_path"`
}

type exampleConfig struct {
	Example string `koanf:"example"`
}

type smtpConfig struct {
	Host    string `koanf:"host"`
	Port    int    `koanf:"port"`
	Domain  string `koanf:"domain"`
	TLSCert string `koanf:"tls_cert"`
	TLSKey  string `koanf:"tls_key"`
}

type smsgateConfig struct {
	URL                 string `koanf:"url"`
	SkipPhoneValidation bool   `koanf:"skip_phone_validation"`
}

type Config struct {
	HTTP    http          `koanf:"http"`
	Example exampleConfig `koanf:"example"`
	SMTP    smtpConfig    `koanf:"smtp"`
	SMSGate smsgateConfig `koanf:"smsgate"`
}

func Default() Config {
	const defaultSMTPPort = 587

	return Config{
		HTTP: http{
			Address:     "127.0.0.1:3000",
			ProxyHeader: "X-Forwarded-For",
			Proxies:     []string{},
			OpenAPI: openAPIConfig{
				Enabled:    true,
				PublicHost: "",
				PublicPath: "",
			},
		},

		Example: exampleConfig{
			Example: "example",
		},

		SMTP: smtpConfig{
			Host:    "127.0.0.1",
			Port:    defaultSMTPPort,
			Domain:  "example.com",
			TLSCert: "",
			TLSKey:  "",
		},

		SMSGate: smsgateConfig{
			URL:                 "https://api.sms-gate.app/3rdparty/v1",
			SkipPhoneValidation: false,
		},
	}
}

func New() (Config, error) {
	cfg := Default()

	options := []config.Option{}
	if yamlPath := os.Getenv("CONFIG_PATH"); yamlPath != "" {
		options = append(options, config.WithLocalYAML(yamlPath))
	}

	if err := config.Load(&cfg, options...); err != nil {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}

	return cfg, nil
}
