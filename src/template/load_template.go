package template

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"projectcreator/util"
)

type Child struct {
	Name     string
	Type     string
	Children []Child
	Content  string
}

type Template struct {
	Name      string
	Tree      []Child
	Gitignore string
	Commands  []struct {
		Command string
		Args    []string
	}
}

func LoadTemplate(templateName string) (Template, error) {
	fileName := fmt.Sprintf("../templates/%v.json", templateName)

	fmt.Println(fileName)

	if !util.IsFileExistent(fileName) {
		return Template{}, errors.New("invalid template")
	}

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	template := Template{}

	if err = json.Unmarshal(data, &template); err != nil {
		log.Fatal("Error parsing JSON file:", err)
	}

	return template, nil
}
