package main

import (
	"fmt"
	"github.com/essentialkaos/translit"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()

	app.Name = "Parser Avito Pages"
	app.Usage = "Can get location, categories and count items in JSON-format"

	var name string
	human := false

	app.Commands = []cli.Command{
		{
			Name:        "get",
			Aliases:     []string{"g"},
			Usage:       "get Options",
			Subcommands: []cli.Command{
				{
					Name:    "locations",
					Aliases: []string{"l"},
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "print, p",
							Destination: &human,
						},
						cli.StringFlag{
							Name:        "name, n",
							Value:       "Moscow",
							Usage:       "Get sublocations of location",
							Destination: &name,
						},
					},
					Subcommands: []cli.Command{
						{
							Name:    "all",
							Aliases: []string{"a"},
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name: "print, p",
									Destination: &human,
								},
							},
							Action: func(c *cli.Context) {
								if human == true {
									printHumanReadableAll()
								} else {
									err := ioutil.WriteFile("locationsTree.json", []byte(getSublocationsJson()), 0644)

									if err != nil {
										log.Fatal(err)
									}
									fmt.Println("Your JSON-file in locationsTree.json")
								}

							},
						},
					},
					Usage: "l [-n] [LOCATION_NAME]",
					Action: func(c *cli.Context) {
						if human == true {
							printHumanReadableOne(name)
						} else {
							fileName := strings.ReplaceAll("locationsTree" + strings.Title(translit.EncodeToBGN(name)), " ", "") + ".json"
							if strJson := getSublocationsOne(name); strJson != `{"message":"No such location"}` {
								err := ioutil.WriteFile(fileName, []byte(strJson), 0644)
								if err != nil {
									log.Fatal(err)
								}
								fmt.Println("Your JSON-file in " + fileName)
							} else {
								fmt.Println("No such location " + name)
							}
						}
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
