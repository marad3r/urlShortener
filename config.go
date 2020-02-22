package main

type Config struct {
	Env    string `yaml:"environment",envconfig:"ENV"`
	Server `yaml:"server"`
}

type Server struct {
	Host         string `yaml:"host",envconfig:"SERVER_HOST"`
	Port         string `yaml:"port",envconfig:"SERVER_PORT"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
	IdleTimeout  int    `yaml:"idleTimeout"`
}
