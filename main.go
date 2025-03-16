package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const colourR = "\033[0;31m"
const colourG = "\033[0;32m"
const colourB = "\033[0;34m"
const colourC = "\033[0;36m"
const colourM = "\033[0;35m"
const colourY = "\033[0;33m"
const colourK = "\033[0;00m"
const emoji_camera = '\U0001F4F7'

func main() {

	// Initial variables
	var yaml_file string
	location := os.Getenv("LOCATION")

	if location == "" {
		log.Fatal("ERROR: missing LOCATION environment variable")
	}

	// Determine argument precedence, cli arguments override environment variables
	if len(os.Args) != 2 {
		if os.Getenv("YAML") == "" {
			log.Fatal("ERROR: missing yaml file argument or YAML environment variable")
		} else {
			yaml_file = os.Getenv("YAML")
		}
	} else if os.Args[1] != "" {
		yaml_file = os.Args[1]
	}

	// Read yaml file into memory
	yamels, err := readYamlEntries(yaml_file)
	if err != nil {
		log.Fatal(err)
	}

	// Display settings before starting
	if os.Getenv("VERBOSE") == "true" {
		log.Printf("LOCATION='%s'\n", location)
		log.Printf("yaml_file='%s'\n", yaml_file)
	}

	// Execute commands
	for _, entry := range yamels.Section {
		if location == entry.Location {
			item := entry.Item
			description := entry.Description
			command := entry.Command
			fmt.Printf("%v%+v %v%+v\n  %vCommand:\n%v%+v%v\n", colourY, item, colourR, description, colourM, colourC, command, colourK)
			cmd := exec.Command("bash", "-c", command)
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal("Command: ", err)
				return
			}
			fmt.Println(string(stdout))
			fmt.Printf("%c\n", emoji_camera)
		}
	}
}
