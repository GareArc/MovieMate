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
	main_config *koanf.Koanf = koanf.New(".")
)

func GetStaticConfig() *koanf.Koanf {
	return main_config
}

func InitConfig() {
	// Load from yaml file
	if err := main_config.Load(file.Provider("env.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Load from EV
	if err := main_config.Load(env.Provider("", ".", func(s string) string {
		return strings.ToLower(strings.ReplaceAll(s, "_", "."))
	}), nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}
