package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

type post struct {
	title       string
	link        string
	upvotes     string
	description string
}
type solution struct {
	description string
	upvotes     string
}

func searchPost(query string) []post {
	res, err := goquery.NewDocument(fmt.Sprintf("https://stackoverflow.com/search?q=%s", strings.ReplaceAll(query, " ", "+")))
	if err != nil {
		log.Fatal(err)
	}

	var items []post
	res.Find(".question-summary").Each(func(index int, item *goquery.Selection) {
		linkTag := item.Find(".result-link").Find("a")
		link, _ := linkTag.Attr("href")
		title := strings.TrimFunc(linkTag.Text(), func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
		description := item.Find(".excerpt").Text()
		upvotes := item.Find(".vote-count-post").Text()
		//fmt.Println(link)
		if true || (strings.HasPrefix(title, "Q:") && index < 4) {
			items = append(items, post{title, link, upvotes, description})
			index++
		}
	})
	return items
}
func getPost(node post) (solution, []solution) {
	urlString := fmt.Sprintf("https://stackoverflow.com/%s", node.link)
	res, err := goquery.NewDocument(urlString)
	if err != nil {
		log.Fatal(err)
	}
	var answers []solution
	question := res.Find(".question").Find(".post-layout")
	answers = append(answers, solution{strings.Trim(question.Find(".post-text").Text(), "\n"), question.Find(".js-vote-count").Text()})

	acceptedContainer := res.Find(".accepted-answer").Find(".post-layout")
	acceptedAnswer := solution{strings.Trim(acceptedContainer.Find(".post-text").Text(), "\n"), acceptedContainer.Find(".js-vote-count").Text()}

	if (acceptedAnswer != solution{}) {
		res.Find(".accepted-answer").NextAll().Each(func(index int, item *goquery.Selection) {

			if item.Find(".post-text").Text() != "" {
				answers = append(answers, solution{strings.Trim(item.Find(".post-text").Text(), "\n"), item.Find(".js-vote-count").Text()})
			}
		})
	} else {
		res.Find(".post-layout").Each(func(index int, item *goquery.Selection) {

			if item.Find(".post-text").Text() != "" {
				answers = append(answers, solution{strings.Trim(item.Find(".post-text").Text(), "\n"), item.Find(".js-vote-count").Text()})
			}
		})
	}

	return acceptedAnswer, answers
}

func getTitleDescription(posts []post) ([]string, []string, []string) {
	var title []string
	var description []string
	var link []string
	for _, temp := range posts {

		title = append(title, strings.TrimFunc(temp.title, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})+" - "+temp.upvotes+" upvotes")
		description = append(description, temp.description)
		link = append(link, temp.link)
	}
	return title, description, link
}
