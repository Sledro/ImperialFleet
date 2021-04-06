package config

import (
	"github.com/spf13/viper"
)

// LoadConfig - Loads the config vars
func LoadConfig() {
	// dbConnString would useually be set as an env var
	viper.SetDefault("dbConnString", "root:password@tcp(127.0.0.1:3306)/ImperialFleet")
}
