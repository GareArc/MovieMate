package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var (
	MainConfig *koanf.Koanf = koanf.New(".")
)

func Init() {
	// Load from yaml file
	if err := MainConfig.Load(file.Provider("env.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Load from EV
	if err := MainConfig.Load(env.Provider("APP_", ".", func(s string) string {
		return strings.ToLower(strings.ReplaceAll(strings.TrimPrefix(s, "APP_"), "_", "."))
	}), nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}
