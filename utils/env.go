package utils

import "github.com/spf13/viper"

type Env struct {
	DBDriver	string `mapstructure:"DB_DRIVER"`
	DBSource	string `mapstructure:"DB_SOURCE"`
	Port			string `mapstructure:"PORT"`
}

func LoadEnv(path string) (env Env, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}
