package config

type Server struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
}
