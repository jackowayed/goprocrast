package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func dev() bool {
	return os.Getenv("DEV") == "1" || true //temp
}

func hostsPath() string {
	if dev() {
		return "./hosts.dev"
	} else {
		return "/etc/hosts"
	}
}

func hostsFileContent() string {
	content, err := ioutil.ReadFile(hostsPath()) // the file is inside the local directory
	if err != nil {
		panic(err)
	}
	return string(content)
}

func activate() {
	fmt.Println(hostsFileContent())
}

func deactivate() {

}

func main() {
	fmt.Println(os.Args[0])
	if len(os.Args) <= 1 {
		fmt.Println("usage")
		os.Exit(0)
	}

	cmd := os.Args[1]
	fmt.Println(cmd)
	switch cmd {
	case "on":
		activate()
	case "off":
		deactivate()
	}
}
