package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"syscall"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	check(err)
	return content
}

func currentHosts() []string {
	file, err := os.Open(".noprocrast")
	check(err)
	scanner := bufio.NewScanner(file)
	var hosts []string
	for scanner.Scan() {
		host := strings.TrimSpace(scanner.Text())
		if len(host) > 0 {
			hosts = append(hosts, host)
		}
	}
	return hosts
}

func activate() {
	//fmt.Println(hostsFileContent())
	fmt.Println(currentHosts())
}

func deactivate() {
	re, err := regexp.Compile("(?m)(\n\n)?# noprocrast start.*# noprocrast end")
	check(err)
	cleanHosts := re.ReplaceAllLiteral(hostsFileContent(), nil)
	fmt.Println(cleanHosts)
}

func suid() {
	err := syscall.Setuid(0)
	check(err)
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
