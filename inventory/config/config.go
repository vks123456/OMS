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

	CacheExpTimeInSeconds time.Duration

	// kafka
	KafkaBrokerUrl string
	KafkaVerbose   bool
	KafkaClientId  string
	KafkaTopic     string
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

		CacheExpTimeInSeconds: viper.GetDuration("CACHE_EXPIRY_TIME_IN_SECONDS"),

		KafkaBrokerUrl: viper.GetString("KAFKA_BROKER_URL"),
		KafkaVerbose:   viper.GetBool("KAFKA_VERBOSE"),
		KafkaClientId:  viper.GetString("KAFKA_CLIENT_ID"),
		KafkaTopic:     viper.GetString("KAFKA_TOPIC"),
	}
}

func setConfigDefaults() {
	viper.SetDefault("APP_NAME", "inventory")
	viper.SetDefault("LISTEN_HOST", "0.0.0.0")
	viper.SetDefault("HTTP_PORT", "8081")

	viper.SetDefault("KAFKA_BROKER_URL", "localhost:19092,localhost:29092,localhost:39092")
	viper.SetDefault("KAFKA_VERBOSE", true)
	viper.SetDefault("KAFKA_CLIENT_ID", "my-kafka-client")
	viper.SetDefault("KAFKA_TOPIC", "order")
}
