package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"syscall"
)

var noprocrastRegexp = regexp.MustCompile(
	"(?m)(?:\n\n)?# noprocrast start(?:\n|.)*# noprocrast end")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dev() bool {
	return os.Getenv("DEV") == "1"
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

func noprocrastPath() string {
	homePath, err := os.UserHomeDir()
	check(err)
	return path.Join(homePath, ".noprocrast")
}

func noprocrastFile() *os.File {
	file, err := os.Open(noprocrastPath())
	check(err)
	return file
}

func currentHosts() []string {
	scanner := bufio.NewScanner(noprocrastFile())
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
		file.WriteString("127.0.0.1 ")
		file.WriteString(host)
		file.WriteString("\n")
	}
	file.WriteString("\n# noprocrast end")
	check(file.Close())

	check(exec.Command("dscacheutil", "-flushcache").Run())
}

func openHostsFile(truncate bool) *os.File {
	var flags int
	if truncate {
		flags = os.O_WRONLY | os.O_TRUNC
	} else {
		flags = os.O_APPEND | os.O_WRONLY
	}
	file, err := os.OpenFile(hostsPath(), flags, 0)
	if err != nil && strings.Contains(err.Error(), "permission denied") &&
		syscall.Getuid() != 0 {
		suidRoot()
		return openHostsFile(truncate)
	} else if err != nil {
		panic(err)
	}
	return file
}

func deactivate() {
	cleanHosts := noprocrastRegexp.ReplaceAllLiteral(hostsFileContent(), nil)
	file := openHostsFile(true)
	file.Write(cleanHosts)
	check(file.Close())
}

func suidRoot() {
	err := syscall.Setuid(0)
	check(err)
}

func edit() {
	editor := os.Getenv("EDITOR")
	if len(editor) == 0 {
		editor = "vi"
	}
	cmd := exec.Command(editor, noprocrastPath())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	check(cmd.Run())
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage")
		os.Exit(0)
	}
	cmd := os.Args[1]
	switch cmd {
	case "on", "reload":
		activate()
	case "off":
		deactivate()
	case "edit":
		edit()
	}
}
