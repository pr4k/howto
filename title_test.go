package main

import (
	"testing"
)

func TestGetTitleDescription(t *testing.T) {

	items := []post{post{"title1", "link1", "10", "description1"}, post{"title2", "link2", "20", "description2"}, post{"title3", "link3", "30", "description3"}}
	title, description, link := getTitleDescription(items)
	expectedTitle := []string{"title1 - 10 upvotes", "title2 - 20 upvotes", "title3 - 30 upvotes"}
	expectedDescription := []string{"description1", "description2", "description3"}
	expectedLinks := []string{"link1", "link2", "link3"}
	if notEqual(title, expectedTitle) || notEqual(description, expectedDescription) || notEqual(link, expectedLinks) {
		t.Error("Error incorrect")
	}

}

func notEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return true
	}
	array := make([]int, len(a))
	for i := range array {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}
