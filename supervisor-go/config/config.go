package config

import "github.com/spf13/viper"

type Program struct {
	Command      string
	GitHubEvents []string
	Name         string
	PID          int
	URL          string
}

func (p *Program) NeedInstall() bool {
	return p.URL != ""
}

type Config struct {
	Github struct {
		Secret string
	}
	HTTP struct {
		Port string
	}
	Programs []Program
}

func Load() (*Config, error) {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.Debug()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
