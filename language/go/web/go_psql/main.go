package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/saurabh3460/go_psql/router"
)

type Dummy struct {
	Name    string  `json:"name"`
	Number  int64   `json:"number"`
	Pointer *string `json:"pointer"`
}

func ReadConfigFile() {
	if err := viper.ReadInConfig(); err != nil {
		if errMsg, ok := err.(viper.ConfigFileNotFoundError); ok {

			panic(fmt.Errorf("%s", errMsg))
		} else if errMsg, ok := err.(*fs.PathError); ok {
			fmt.Printf("File used %s\n", viper.ConfigFileUsed())
			panic(fmt.Errorf("config path not found %s", errMsg))
		} else {
			fmt.Printf("%T\n", err)
			panic(fmt.Errorf("fetal error config file %s", err))
		}
	} else {
		fmt.Println("Config file loaded successfully...")
	}
}
func loadEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./config")
}
func main() {
	loadEnv()
	ReadConfigFile()
	// DB := viper.GetString("DATABASE") // add error handling here

	router := router.Router()
	fmt.Println("Starting go server on 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
