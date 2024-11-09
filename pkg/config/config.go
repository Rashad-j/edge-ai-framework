package config

type Config struct {
	ModelPath string `json:"model_path"`
	Port      int    `json:"port"`
}

// InitConfig get the env configutation
func InitConfig() *Config {
	return &Config{}
}
