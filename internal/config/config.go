package config

import "github.com/spf13/viper"

type Config struct {
    AppPort   string
    DBUrl     string
    JWTSecret string
}

func LoadConfig() (*Config, error) {
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    return &Config{
        AppPort:   viper.GetString("APP_PORT"),
        DBUrl:     viper.GetString("DB_URL"),
        JWTSecret: viper.GetString("JWT_SECRET"),
    }, nil
}
