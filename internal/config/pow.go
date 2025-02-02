package config

type PoW struct {
	Difficulty int64 `yaml:"difficulty"`
	TokenSize  int64 `yaml:"token-size"`
}
