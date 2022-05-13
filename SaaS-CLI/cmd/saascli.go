package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zackmacharia/PANOS-GOLANG/SaaS-CLI/pkg/key"
	"github.com/zackmacharia/PANOS-GOLANG/SaaS-CLI/pkg/saasreport"
)

var app = cli.NewApp()

func info() {
	app.Name = "Simple PANOS SaaS CLI App"
	app.Usage = "Displays SaaS applications on firewall and add Sanctioned tag to applications"
	app.Authors = []*cli.Author{
		{
			Name:  "Zack Macharia",
			Email: "zmacharia@paloaltonetworks.com",
		},
	}
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []*cli.Command{

		{
			Name:    "genkey",
			Aliases: []string{"key"},
			Usage:   "Generates API Key",
			Action: func(c *cli.Context) error {
				apikey.GetApiKey()
				return nil
			},
		},
		{
			Name:    "displaySaaSApps",
			Aliases: []string{"dsa"},
			Usage:   "Displays SaaS Applications",
			Action: func(c *cli.Context) error {
				saasreport.DisplaySaaSReport()
				return nil
			},
		},
		{
			Name:    "writesaasfile",
			Aliases: []string{"wsf"},
			Usage:   "Creates a file with the SaaS application names to a file",
			Action: func(c *cli.Context) error {
				data, err := saasreport.PullSaaSReport()
				if err != nil {
					log.Fatal(err)
				}
				saasreport.CreateSaaSAppsFile(data)
				return nil
			},
		},
		{
			Name:    "addSanctionedTag",
			Aliases: []string{"ast"},
			Usage:   "Add a Sanctioned tag to an application",
			Action: func(c *cli.Context) error {
				saasreport.AddSanctionedTag()
				return nil
			},
		},
	}
}

func main() {

	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
