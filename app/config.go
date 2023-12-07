package app

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewConfig() {
	viper.SetConfigName("config")        // name of config file (without extension)
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/aibook/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.aibook") // call multiple times to add many search paths
	viper.AddConfigPath(".")             // optionally look for config in the working directory
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
