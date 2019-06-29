package main

import (
	"fmt"
	"github.com/essentialkaos/translit"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

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
			Usage:       "get [\"locations (shortly l), categories (c), statistics (s)\"]",
			Subcommands: []cli.Command{
				{
					Name:    "locations",
					Aliases: []string{"l"},
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:        "print, p",
							Usage: " - print in human readable format",
							Destination: &human,
						},
						cli.StringFlag{
							Name:        "name, n",
							Value:       "Москва",
							Usage:       "get sublocations of location",
							Destination: &name,
						},
					},
					Subcommands: []cli.Command{
						{
							Name:    "all",
							Aliases: []string{"a"},
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name:        "print, p",
									Usage: " - print in human readable format",
									Destination: &human,
								},
							},
							Action: func(c *cli.Context) {
								if human == true {
									err := printHumanReadableAll()

									check(err)
								} else {
									err := ioutil.WriteFile("/go/data/locationsTree.json", []byte(getSublocationsJson()), 0644)
									check(err)
									fmt.Println("Your JSON-file in ./data/locationsTree.json")
								}

							},
						},
					},
					Usage: "with flag -n enter location in russian",
					Action: func(c *cli.Context) {
						if human == true {
							err := printHumanReadableOne(name)

							check(err)
						} else {
							fileName := strings.ReplaceAll("/go/data/locationsTree"+strings.Title(translit.EncodeToBGN(name)), " ", "") + ".json"
							if strJson := getSublocationsOne(name); strJson != `{"message":"No such location"}` {
								err := ioutil.WriteFile(fileName, []byte(strJson), 0644)
								check(err)
								fmt.Println("Your JSON-file in ./data/" + strings.Title(translit.EncodeToBGN(name)))
							} else {
								fmt.Println("No such location " + name)
							}
						}
					},
				},
				{
					Name:    "categories",
					Aliases: []string{"c"},
					Usage: "get tree of categories",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:        "print, p",
							Usage: " - print in human readable format",
							Destination: &human,
						},
					},
					Action: func(c *cli.Context) {
						if human == true {
							err := printHumanReadableCategories()
							check(err)
						} else {
							categories, err := getAllCategoriesJson()
							err = ioutil.WriteFile("/go/data/categories.json", categories, 0644)
							check(err)
							if err == nil {
								fmt.Println("Your JSON-file in ./data/categories.json")
							}
						}


					},
				},
				{
					Name:    "statistics",
					Aliases: []string{"s"},
					Usage: "get statistics for locations",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "name, n",
							Destination: &name,
						},
						cli.BoolFlag{
							Name: "print, p",
							Usage: " - print in human readable format",
							Destination: &human,
						},
					},
					Subcommands: []cli.Command{
						{
							Name: "all",
							Aliases: []string{"a"},
							Usage: "get total statistics",
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name:        "print, p",
									Usage: " - print in human readable format",
									Destination: &human,
								},
							},
							Action: func(c *cli.Context) {
								if human == true {
									err := printHumanReadableStatistic("Россия")
									check(err)
								} else {
									stats, err := getStatisticsOneJson("Россия")
									if err != nil {
										check(err)
										return
									}
									fileName := "/go/data/statisticsTotal.json"
									err = ioutil.WriteFile(fileName, stats, 0644)
									check(err)
									if err == nil {
										fmt.Println("Your JSON-file in " + fileName)
									}
								}
							},
						},
					},
					Action: func(c *cli.Context) {
						if human == true {
							err := printHumanReadableStatistic(name)
							check(err)
						} else {
							stats, err := getStatisticsOneJson(name)
							if err != nil {
								check(err)
								return
							}
							fileName := "/go/data/statistics" + strings.Title(translit.EncodeToBGN(name)) + ".json"
							err = ioutil.WriteFile(fileName, stats, 0644)
							check(err)
							if err == nil {
								fmt.Println("Your JSON-file in ./data/statistics" + strings.Title(translit.EncodeToBGN(name)) + ".json")
							}
						}
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	check(err)
}
