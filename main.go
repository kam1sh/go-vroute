package main

import (
	"github.com/spf13/viper"
	"os"
	"log"
	"errors"
)

func check(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, "-", err)
	}
}

func main() {
	app := GetApp()
	err := app.Run(os.Args)
	check(err, "Error")

}

func loadConfigFromVariable(name string) (*viper.Viper, error) {
	configPath := os.Getenv(name)
	if configPath == "" {
		return nil, errors.New("No config provided via environment variable")
	}

	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	cfg := viper.New()
	cfg.SetConfigType("yaml")

	err = cfg.ReadConfig(configFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
