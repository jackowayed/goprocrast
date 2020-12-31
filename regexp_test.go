package main

import (
	"io/ioutil"
	"regexp"
	"testing"
)

// debugging why our regexp isn't matching

func testFileContent() []byte {
	c, err := ioutil.ReadFile("regexptest.txt")
	check(err)
	return c
}

func TestBasicMultiline(t *testing.T) {
	re := regexp.MustCompile("(?m)hello\nworld")
	idxs := re.FindIndex(testFileContent())
	if len(idxs) == 0 {
		t.Error("no match")
	}
}

func TestNoprocrastRegexp(t *testing.T) {
	idxs := noprocrastRegexp.FindIndex(testFileContent())
	if len(idxs) == 0 {
		t.Error("no match")
	}
}
