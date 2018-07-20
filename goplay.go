package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

var filename = "playground.go"

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dir := usr.HomeDir
	dir = filepath.Join(dir, "Documents")

	err = os.Chdir(dir)
	if err != nil {
		log.Println(err)
	}

	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	cmd := exec.Command("vim", "playground.go")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
	cmd.Run()
}
