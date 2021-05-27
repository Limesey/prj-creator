package util

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Settings struct {
	Defaults struct {
		ProjectDir string
		Git        bool
		Github     bool
	}
}

func CreateSettings() Settings {
	return Settings{}
}

func GetSettings() (Settings, error) {
	const FILE string = "../settings.json"

	if !IsFileExistent(FILE) {
		return Settings{}, errors.New("no settings file")
	} else {
		data, err := os.ReadFile(FILE)

		if err != nil {
			log.Fatal("Unable to read file: ", err)
		}

		settings := Settings{}

		err = json.Unmarshal(data, &settings)

		if err != nil {
			log.Fatal("Unmarshall error: ", err)
		}

		return settings, nil
	}
}
