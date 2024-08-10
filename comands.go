package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func createSubs(root string, folder *Folder) {
	root = strings.ReplaceAll(root, "\\", "/")
	subRoot := root + "/" + folder.name
	err := os.Mkdir(subRoot, 0777)
	if err != nil {
		fmt.Println("Error creating file: ", err.Error())
		return
	}

	for _, subFolder := range folder.children {
		createSubs(subRoot, &subFolder)
	}
}

func initProject(c *cli.Context) error {
	fmt.Println("Starting to create folder structure")
	var project string

	if project = c.String("project-name"); project == "" {
		fmt.Println("Please provide the project name: ")
		fmt.Scan(&project)
	}

	root, err := os.Getwd()
	if err != nil {

		return fmt.Errorf("problems creating initial project strutuctre: %s", err.Error())
	}
	createSubs(root, &Folder{project, initialFileStruct})
	return nil
}

func addModule(c *cli.Context) (err error) {
	var root string
	var moduleName string

	if root = c.String("root"); root == "" {
		root, err = os.Getwd()
		if err != nil {
			log.Fatal("Couldnt get current working directory: ", err.Error())
		}
	}
	root = strings.ReplaceAll(root, "\\", "/")

	if moduleName = c.String("name"); moduleName == "" {
		fmt.Println("Enter the name of the module: ")
		fmt.Scan(&moduleName)
	}

	currentDir, _ := os.Getwd()
	currentDir = strings.ReplaceAll(currentDir, "\\", "/")

	if root != currentDir {
		currentDir += "/" + root
	}
	currentDir += "/_Modules"

	folder := &Folder{
		name:     moduleName,
		children: BaseModulesFolder,
	}

	createSubs(currentDir, folder)

	return nil
}
