package main

import (
	"fmt"
	"log"
  "os"
	"os/exec"
)

func main(){
  yamels,err := readYamlEntries(os.Args[1])
   if err != nil {
    log.Fatal(err)
  }

  for _, entry := range yamels.Section {
    item := entry.Item
    command := entry.Command
    fmt.Printf("ITEM: %+v\n  Command:\n%+v\n",item,command)
    cmd := exec.Command("bash", "-c", command)
    stdout, err := cmd.Output()
    if err != nil {
      log.Fatal(err)
      return
    }
    fmt.Println(string(stdout))
  }
}

