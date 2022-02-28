package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

/*
░░░░░░░ ░░ ░░      ░░░░░░░     ░░   ░░  ░░░░░  ░░░    ░░ ░░░░░░  ░░      ░░ ░░░    ░░  ░░░░░░
▒▒      ▒▒ ▒▒      ▒▒          ▒▒   ▒▒ ▒▒   ▒▒ ▒▒▒▒   ▒▒ ▒▒   ▒▒ ▒▒      ▒▒ ▒▒▒▒   ▒▒ ▒▒
▒▒▒▒▒   ▒▒ ▒▒      ▒▒▒▒▒       ▒▒▒▒▒▒▒ ▒▒▒▒▒▒▒ ▒▒ ▒▒  ▒▒ ▒▒   ▒▒ ▒▒      ▒▒ ▒▒ ▒▒  ▒▒ ▒▒   ▒▒▒
▓▓      ▓▓ ▓▓      ▓▓          ▓▓   ▓▓ ▓▓   ▓▓ ▓▓  ▓▓ ▓▓ ▓▓   ▓▓ ▓▓      ▓▓ ▓▓  ▓▓ ▓▓ ▓▓    ▓▓
██      ██ ███████ ███████     ██   ██ ██   ██ ██   ████ ██████  ███████ ██ ██   ████  ██████
*/

type FileHandler struct {
	prefsFile string
	dataFile  string
}

type Entry struct {
	Id   int
	Text string
}

// ▌║█║▌│║▌│║▌║▌█║ List Data ▌│║▌║▌│║║▌█║▌║█

func (f *FileHandler) SaveList(data json.RawMessage) (string, error) {

	/*****
	* The web portion handles the to do list entries as JSON,
	* which is what will be sent here for saving.  So we need
	* to unmarshal the JSON so it's in a format Go can deal with.
	 */

	var items []Entry
	json.Unmarshal([]byte(data), &items)

	fmt.Sprintf("Saver variables\ndata: %v+\nitems: %v+", data, items)

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fileName := filepath.Join(cwd, "bees!")
	dataFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating data file")
		return "", err
	}
	_, err = dataFile.Write(data)
	if err != nil {
		fmt.Println("Error writing to data file")
		return "", err
	}
	fmt.Println("Data file written successfully")
	err = dataFile.Close()
	if err != nil {
		fmt.Println("Couldn't close data file")
		return "", err
	}

	msg := fmt.Sprintf("Data saved: %v+", data)
	return msg, nil

}

func (f *FileHandler) LoadList() (string, error) {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory")
		return "", err
	}

	var rawData []byte

	fileName := filepath.Join(cwd, "bees!")
	rawData, err = os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Sprintf("Got data:\n%v+", rawData)
	return string(rawData), nil

}
