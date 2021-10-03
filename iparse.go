package main

import (
	"fmt"
	"log"
	"math/rand"

	"time"

	"github.com/zhengyi13/iparse/repos"
)

func main() {
	r, _ := repos.GetRepos()
	rand.Seed(time.Now().Unix())

	fmt.Println("Here are all the configured repos: ")
	for _, repo := range r {
		fmt.Printf("%s\n", repo)
	}

	fmt.Println("Here is a random repo's data: ")

	selectedRepo := r[rand.Intn(len(r))]

	rd, err := repos.GetRepoIni(selectedRepo)

	if err != nil {
		log.Fatalf("Died getting ini info: %v", err)
	}

	fmt.Printf("\n\n %s", string(rd))

}
