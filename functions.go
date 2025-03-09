package main

import (
	"os"

	"github.com/goccy/go-yaml"
)

type yamlEntriesOne struct {
  Section []struct {
    Item        string `yaml:"item"`
    Description string `yaml:"description"`
    Notes       string `yaml:"notes"`
    Location    string `yaml:"location"`
    Command     string `yaml:"command"`
  } `yaml:"section"`
}

type yamlEntries struct {
  Section []Section `yaml:"section"`
}

type Section struct {
    Item        string `yaml:"item"`
    Description string `yaml:"description"`
    Notes       string `yaml:"notes"`
    Location    string `yaml:"location"`
    Command     string `yaml:"command"`
}

func readYamlEntries(s string) (yamlEntries, error) {

  // Load the file, returns []byte
  f, err := os.ReadFile(s)
  if err != nil {
    return yamlEntries{}, err
  }

  // Create empty item to be target of Unmarshalling
  var c yamlEntries 

  // Unmarshal the input YAML file into empty itemStructure (var c)
  if err := yaml.Unmarshal(f, &c); err != nil {
    return yamlEntries{}, err
  }
  //fmt.Printf("%+v\n",c)

  return c, nil
}
