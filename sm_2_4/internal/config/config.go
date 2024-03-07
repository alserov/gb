package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Env     string `yaml:"env"`
	Port    int    `yaml:"port"`
	Timeout struct {
		Read  time.Duration `yaml:"read"`
		Write time.Duration `yaml:"write"`
	} `yaml:"timeout"`
	DB struct {
		Addr string `yaml:"addr"`
	}
}

func MustLoad() *Config {
	path := fetchPath()
	if path == "" {
		panic("path to cfg file must be provided")
	}

	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read cfg file: " + err.Error())
	}

	var cfg Config
	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to parse config file: " + err.Error())
	}

	return &cfg
}

func fetchPath() string {
	var path string
	flag.StringVar(&path, "c", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CFG_PATH")
	}

	return path
}
