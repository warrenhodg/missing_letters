package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var words = []string{
	"__i__",
	"V__lat__n",
	"H__dw__e",
	"__rses__e",
	"P__sev__e",
	"S__tim__t",
	"__da__ted",
	"C__kb__k",
	"Lo__i__",
	"D__tingu__h",
	"P__dl__",
	"S__ur__e",
	"__p__zard",
	"Who__sa__",
	"__at__",
	"__od__rk",
	"__ma__",
	"__y__rd",
	"H__rtbr__k",
	"B__evol__t",
	"C__diti__",
	"An__cipa__on",
	"__rri__lum",
	"E__agi__",
	"Inc__p__ate",
}

var allWords []string

func match(p string, s string, o1 int, o2 int) bool {
	if len(p) != len(s) {
		return false
	}
	for i := range p {
		c := p[i]
		if c == '_' {
			continue
		}
		if s[i] != c {
			return false
		}
	}
	if s[o1] != s[o2] || s[o1+1] != s[o2+1] {
		return false
	}
	return true
}

func findWord(p string) string {
	o1 := strings.Index(p, "_")
	o2 := strings.Index(p[o1+2:], "_") + o1 + 2

	for _, s := range allWords {
		if match(p, s, o1, o2) {
			return s
		}
	}

	return ""
}

func main() {
	dict, _ := ioutil.ReadFile("./words.txt")
	allWords = strings.Split(string(dict), "\n")

	for _, p := range words {
		fmt.Printf("%s : %s\n", p, findWord(strings.ToLower(p)))
	}
}
