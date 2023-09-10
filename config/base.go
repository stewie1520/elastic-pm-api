package config

import "github.com/spf13/viper"

type Config struct {
	Port          int16  `mapstructure:"PORT"`
	ApiDomain     string `mapstructure:"API_DOMAIN"`
	WebsiteDomain string `mapstructure:"WEBSITE_DOMAIN"`
	SuperTokens   struct {
		ConnectionUrl string `mapstructure:"CONNECTION_URL"`
		ApiKey        string `mapstructure:"API_KEY"`
	} `mapstructure:"SUPERTOKENS"`
	GinMode string `mapstructure:"GIN_MODE"`
}

func Init() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
