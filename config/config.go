package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
    RabbitUser string
    RabbitPassword string
    RabbitHost string
    RabbitPort string
    ExchangeName string

    CassandraHost string
    CassandraPort int
}

var (
    config *Config
)

func LoadConfig() error {
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    config = &Config{}

    neededVars := []string{
        "RABBIT_USER",
        "RABBIT_PASSWORD",
        "RABBIT_HOST",
        "RABBIT_PORT",
        "EXCHANGE_NAME",
        "CASSANDRA_HOST",
        "CASSANDRA_PORT",
    }
    for _, variable := range neededVars {
        if ok := viper.IsSet(variable); !ok {
            return fmt.Errorf("the variable %s does not exist", variable)
        }
    }
    config.RabbitUser = viper.GetString("RABBIT_USER")
    config.RabbitPassword = viper.GetString("RABBIT_PASSWORD")
    config.RabbitHost = viper.GetString("RABBIT_HOST")
    config.RabbitPort = viper.GetString("RABBIT_PORT")
    config.ExchangeName = viper.GetString("EXCHANGE_NAME")
    config.CassandraHost = viper.GetString("CASSANDRA_HOST")
    config.CassandraPort = viper.GetInt("CASSANDRA_PORT")
    return nil
}

func GetConfig() *Config {
    return config
}
