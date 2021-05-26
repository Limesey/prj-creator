package project

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"projectcreator/template"
)

func createChild(child template.Child, dir string) {
	if child.Type == "directory" {
		dir = fmt.Sprintf("%v/%v", dir, child.Name)

		err := os.MkdirAll(dir, os.ModePerm)

		if err != nil {
			log.Fatal("Error creating directory:", err)
		}

		for i := range child.Children {
			createChild(child.Children[i], dir)
		}

	} else if child.Type == "file" {
		file, err := os.Create(fmt.Sprintf("%v/%v", dir, child.Name))

		if err != nil {
			log.Fatal("Error creating file:", err)
		}

		err = os.WriteFile(file.Name(), []byte(child.Content), os.ModePerm)

		if err != nil {
			log.Fatal("Error writing file:", err)
		}

	} else {
		fmt.Printf("Unable to create child '%v' of type '%v': Unknown type", child.Name, child.Type)
	}
}

func Create(template template.Template, dir string, name string) {
	dir = fmt.Sprintf("%v/%v", dir, name)

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal("Error creating project directory:", err)
	}

	for i := range template.Tree {
		child := template.Tree[i]

		createChild(child, dir)
	}

	absPath, err := filepath.Abs(dir)

	if err != nil {
		log.Fatal("Error getting absolute path:", err)
	}

	err = os.Chdir(absPath)

	if err != nil {
		fmt.Println("Unable to switch directories:", err)
	}

	for i := range template.Commands {
		templateCommand := template.Commands[i]
		command := templateCommand.Command
		args := templateCommand.Args

		cmd := exec.Command(command)
		cmd.Args = make([]string, len(args)+1)

		for j := range args {
			cmd.Args[j+1] = args[j]
		}

		output, err := cmd.Output()

		if err != nil {
			log.Fatal("Error during the execution of command ", templateCommand.Command, ":", err)
		}

		fmt.Println(string(output))
	}
}
