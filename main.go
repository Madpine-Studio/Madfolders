package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
	"time"
)

type Folder struct {
	name     string
	children []Folder
}

func main() {
	app := &cli.App{
		Name: "Madfolders",
		Usage: "CLI tool to help studio project structure, providing easy " +
			"to use helper function to iterate over the project",
		Version:  "0.0.1",
		Compiled: time.Time{},
		Authors: []cli.Author{
			{
				Name:  "Jaw",
				Email: "tuyweber@gmail.com",
			},
			{
				Name:  "Murilo",
				Email: "murilochayel@gmail.com",
			},
			{
				Name:  "Schermak",
				Email: "insert schermak email :D",
			},
		},
		Copyright: "Madpinestudio",
		Commands: []cli.Command{
			{
				Name:   "init",
				Usage:  "Creates the initial madpine project structure",
				Action: initProject,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "project-name",
						Usage:    "Provides the project name instead of completing it on scanf statement",
						Required: false,
					},
				},
			},
			{
				Name:  "add-module",
				Usage: "Add a module component to project including subfolders structured",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: "root",
						Usage: func() string {
							currentDir, _ := os.Getwd()
							currentDir = strings.ReplaceAll(currentDir, "\\", "/")
							return fmt.Sprintf("The root directory of project DEFAULT[%s]", currentDir)
						}(),
						Required: false,
					},
					cli.StringFlag{
						Name:     "name",
						Usage:    "Name of the module",
						Required: false,
					},
				},
				Action: addModule,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Couldnt start the CLI due to:", err)
	}
}
