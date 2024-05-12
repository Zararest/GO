package config

type LoggerConfig struct {
	Level uint
}

type ServerConfig struct {
	// FIXME
}

type GlobalConfig struct {
	log LoggerConfig
	server ServerConfig
}

func GetParameters() (GlobalConfig, error) {
	// FIXME
	return GlobalConfig{LoggerConfig{0}, ServerConfig{}}, nil
}
