package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Couldnt start the CLI due to:", err)
	}
}
