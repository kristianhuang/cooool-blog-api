package config

import "blog-go/internal/apiserver/options"

type Config struct {
	*options.Options
}

func NewConfig(options *options.Options) *Config {
	return &Config{Options: options}
}
