package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func createSubs(root string, folder *Folder) {
	root = strings.ReplaceAll(root, "\\", "/")
	subRoot := root + "/" + folder.name
	err := os.Mkdir(subRoot, 0777)
	if err != nil {
		return
	}

	for _, subFolder := range folder.children {
		createSubs(subRoot, &subFolder)
	}
}

func initProject(c *cli.Context) error {
	fmt.Println("Starting to create folder structure")
	var project string

	fmt.Println("Please provide the project name: ")
	fmt.Scanln(&project)

	root, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("problems creating initial project strutuctre: ", err.Error())
	}
	createSubs(root, &Folder{project, initialFileStruct})
	return nil
}
