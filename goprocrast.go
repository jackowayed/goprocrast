package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"syscall"
)

func dev() bool {
	return os.Getenv("DEV") == "1" || true //temp
}

func hostsPath() string {
	if dev() {
		return "./hosts.dev"
	}
	return "/etc/hosts"
}

func hostsFileContent() []byte {
	content, err := ioutil.ReadFile(hostsPath())
	if err != nil {
		panic(err)
	}
	return content
}

func activate() {
	fmt.Println(hostsFileContent())
}

func deactivate() {
	re, err := regexp.Compile("(?m)(\n\n)?# noprocrast start.*# noprocrast end")
	if err != nil {
		panic(err)
	}
	cleanHosts := re.ReplaceAllLiteral(hostsFileContent(), nil)
	fmt.Println(cleanHosts)
}

func suid() {
	err := syscall.Setuid(0)
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage")
		os.Exit(0)
	}
	//suid()

	cmd := os.Args[1]
	fmt.Println(cmd)
	switch cmd {
	case "on":
		activate()
	case "off":
		deactivate()
	}
}
