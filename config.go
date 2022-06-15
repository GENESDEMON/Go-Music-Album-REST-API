package main

import (
	"log"

	"github.com/spf13/viper"
)

//Creating a structure that contains our port number and mysql connection string
type Config struct {
	Port          string `mapstructure:"port"`
	ConnectString string `mapstructure:"connection_string"`
}

var AppConfig *Config //App config variable needed by other files

func ConfigLoader() {
	log.Println("Loading server Configurations....") //Print config connections
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
