package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"stage":    "development",
		"version":  0.3,
		"host":     3000,
		"database": "dev",
		"username": "shanu",
	}
)

// https://gist.github.com/saurabh3460/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
func main() {
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// p, _ := os.Getwd()
	// viper.AddConfigPath(path.Join(p, "config"))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if errMsg, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("%s", errMsg))
		} else {
			fmt.Printf("%T\n", err)
			panic(fmt.Errorf("Fetal error config file: %s\n", err))
		}

	}

	fmt.Println(viper.GetBool("username"))
}
