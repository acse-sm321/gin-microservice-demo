// Dive deep in unit test for golang:https://semaphoreci.com/blog/unit-testing
// The unit test for articles util functions
package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// then check the length of the returned list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check if each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}

// for further tests
// func TestXxx(t *testing.T) {

// }
