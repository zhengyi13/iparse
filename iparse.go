package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func getRepos() (repos []string) {
	repolist, err := exec.Command("/usr/bin/dnf", "repolist").Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(repolist), "\n")
	var rs []string
	for _, line := range lines {
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}
		reponame := words[0]
		if reponame == "repo" {
			continue
		}
		rs = append(rs, reponame)
	}
	return rs
}

func main() {
	r := getRepos()
	for _, repo := range r {
		fmt.Printf("%s\n", repo)
	}
}
