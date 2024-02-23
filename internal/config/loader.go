package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

func loadString(name string) string {
	isExist(name)
	return viper.GetString(name)
}

func loadBool(name string) bool {
	isExist(name)
	return viper.GetBool(name)
}

func loadDuration(name string) time.Duration {
	isExist(name)
	return viper.GetDuration(name)
}

func loadInt(name string) int {
	isExist(name)
	return viper.GetInt(name)
}

func isExist(name string) {
	exists := viper.IsSet(name)
	if !exists {
		log.Printf("the vairable [%s] is not set\n", name)
		os.Exit(1)
	}
}
