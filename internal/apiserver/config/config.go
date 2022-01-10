package config

import "blog-api/internal/apiserver/options"

type Config struct {
	*options.Options
}

func NewConfig(options *options.Options) *Config {
	return &Config{Options: options}
}
