package main

import (
	"log"

	"github.com/kam1sh/go-vroute/linux"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var version = "0.1.0"
var cfg *viper.Viper

// GetApp builds CLI Application object
func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "vroute"
	app.Usage = "VPN route management CLI"
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:   "load",
			Usage:  "Load networks from file",
			Action: load,
			Before: prepareApp,
		},
	}
	return app
}

func prepareApp(c *cli.Context) error {
	var err error
	cfg, err = loadConfigFromVariable("VROUTE_CONFIG")
	if err != nil {
		return err
	}

	link := cfg.GetString("vpn.route_to.interface")
	if link == "" {
		log.Fatalln("Interface not provided")
	}
	linuxRouter, err := linux.GetRouter(
		link,
		cfg.GetInt("vpn.table_id"),
		cfg.GetInt("vpn.rule.priority"),
	)
	check(err, "Error building router")

	log.Println(linuxRouter) // TODO

	return err
}

func load(c *cli.Context) error {
	_ = c.Args().First()
	return nil
}
