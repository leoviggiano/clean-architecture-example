package config

import (
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

const (
	Production  = "production"
	Development = "development"
	Test        = "test"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}

	viper.AutomaticEnv()

	viper.SetDefault("max_idle_connections", 50)
	viper.SetDefault("max_open_connections", 100)
	viper.SetDefault("max_lifetime_connections", 60)
}

func Environment() string {
	switch viper.GetString("environment") {
	case "development":
		return Development
	case "test":
		return Test
	default:
		return Production
	}
}

func ServerPort() string {
	return ":4040"
}

func QueryLog() bool {
	return viper.GetBool("query_log")
}

func RedisAddress() string {
	return viper.GetString("redis_address")
}

func RedisDB() int {
	return viper.GetInt("redis_db")
}

func DatabaseConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_name"),
	)
}

func MaxLifetimeConnections() time.Duration {
	maxLifetimeConnections := viper.GetInt("max_lifetime_connections")
	return time.Duration(maxLifetimeConnections * int(time.Minute))
}

func MaxIdleConnections() int {
	return viper.GetInt("max_idle_connections")
}

func MaxOpenConnections() int {
	return viper.GetInt("max_open_connections")
}
