// file: list_posts.go
package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("StackFlow", "Searches on Stack Overflow and return Answers")

	s := parser.String("q", "query", &argparse.Options{Required: true, Help: "Query to Search"})
	// Parse input

	query := *s

	query = strings.Replace(query, "\n", "", -1)
	if query == "" && len(os.Args) > 1 {
		query = "How to " + strings.Join(os.Args[1:], " ")
		search(query)
	} else if query == "" && len(os.Args) == 1 {
		err := parser.Parse(os.Args)
		if err != nil {

			fmt.Print(parser.Usage(err))
		}

	} else {
		query = "How to " + query
		search(query)
	}
}
func search(query string) {
	fmt.Println("Searching query ", query)
	items := searchPost(query)
	var wg sync.WaitGroup

	wg.Add(len(items))
	messages := make(chan string, len(items))

	title, description, _ := getTitleDescription(items)
	acceptedAnswers := make(map[int]solution)
	allAnswers := make(map[int][]solution)
	gettingPost := func(link string, index int, item post) {

		defer wg.Done()
		acceptedTemp, temp := getPost(item)
		allAnswers[index] = temp
		acceptedAnswers[index] = acceptedTemp

		messages <- "done"
	}

	for index, item := range items {

		go gettingPost(item.link, index, item)

	}

	wg.Wait()

	createUI(title, description, allAnswers, acceptedAnswers)
}
