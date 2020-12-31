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

var noprocrastRegexp = regexp.MustCompile("(?m)(\n\n)?# noprocrast start.*# noprocrast end")

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
	deactivate()
	file, err := os.OpenFile(hostsPath(), os.O_APPEND|os.O_WRONLY, 0)
	check(err)
	file.WriteString("\n\n# noprocrast start\n")
	for _, host := range currentHosts() {
		file.WriteString(host)
		file.WriteString("\n")
	}
	file.WriteString("\n# noprocrast end")
	check(file.Close())
}

func deactivate() {
	//fmt.Println(re.FindIndex(hostsFileContent()))
	cleanHosts := noprocrastRegexp.ReplaceAllLiteral(hostsFileContent(), nil)
	file, err := os.OpenFile(hostsPath(), os.O_WRONLY, 0)
	check(err)
	file.Write(cleanHosts)
	check(file.Close())
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
