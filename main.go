package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

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
			command := entry.Command
			fmt.Printf("ITEM: %+v\n  Command:\n%+v\n", item, command)
			cmd := exec.Command("bash", "-c", command)
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println(string(stdout))
		}
	}
}
