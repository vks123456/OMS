package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	//Application Configs
	AppName  string
	Host     string
	HTTPPort string

	// kafka
	KafkaBrokerUrl string
	KafkaVerbose   bool
	KafkaClientId  string
	KafkaTopic     string

	InventoryUrl string

	//Redis Configs
	RedisAddress        string
	RedisPassword       string
	RedisDB             int
	RedisDialTimeoutMs  time.Duration
	RedisReadTimeoutMs  time.Duration
	RedisWriteTimeoutMs time.Duration
}

/*
Initialize all configs and set defaults for them.
*/
func InitConfigStore() *Config {
	viper.AutomaticEnv()
	setConfigDefaults()

	return &Config{
		//Application Configs
		AppName:  viper.GetString("APP_NAME"),
		Host:     viper.GetString("LISTEN_HOST"),
		HTTPPort: viper.GetString("HTTP_PORT"),

		KafkaBrokerUrl: viper.GetString("KAFKA_BROKER_URL"),
		KafkaVerbose:   viper.GetBool("KAFKA_VERBOSE"),
		KafkaClientId:  viper.GetString("KAFKA_CLIENT_ID"),
		KafkaTopic:     viper.GetString("KAFKA_TOPIC"),

		InventoryUrl: viper.GetString("INVENTORY_BASE_URL"),

		//Redis Configs
		RedisAddress:        viper.GetString("REDIS_ADDRESS"),
		RedisPassword:       viper.GetString("REDIS_PASSWORD"),
		RedisDB:             viper.GetInt("REDIS_DB"),
		RedisDialTimeoutMs:  viper.GetDuration("REDIS_DIAL_TIMEOUT_MS)"),
		RedisReadTimeoutMs:  viper.GetDuration("REDIS_READ_TIMEOUT_MS"),
		RedisWriteTimeoutMs: viper.GetDuration("REDIS_WRITE_TIMEOUT_MS"),
	}
}

func setConfigDefaults() {
	viper.SetDefault("APP_NAME", "cart")
	viper.SetDefault("LISTEN_HOST", "0.0.0.0")
	viper.SetDefault("HTTP_PORT", "8080")

	viper.SetDefault("KAFKA_BROKER_URL", "localhost:19092,localhost:29092,localhost:39092")
	viper.SetDefault("KAFKA_VERBOSE", true)
	viper.SetDefault("KAFKA_CLIENT_ID", "my-kafka-client")
	viper.SetDefault("KAFKA_TOPIC", "order")
	viper.SetDefault("INVENTORY_BASE_URL", "http://0.0.0.0:8081/getQuantity")

	viper.SetDefault("REDIS_ADDRESS", "localhost:6379")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("REDIS_DIAL_TIMEOUT_MS", 1000)
	viper.SetDefault("REDIS_READ_TIMEOUT_MS", 1000)
	viper.SetDefault("REDIS_WRITE_TIMEOUT_MS", 1000)
}
