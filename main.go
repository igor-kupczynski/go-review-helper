package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/igor-kupczynski/ghtoken"
	"golang.org/x/oauth2"
)

type byChanges []*github.CommitFile

func (s byChanges) Len() int {
	return len(s)
}
func (s byChanges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byChanges) Less(i, j int) bool {
	return *s[i].Changes < *s[j].Changes
}

func main() {
	ctx := context.Background()

	org, repo := os.Args[1], os.Args[2]
	pr, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("review-helper: %v\n", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("review-helper: %v\n", err)
	}

	t, err := ghtoken.EnsureToken(
		fmt.Sprintf("%s/.help-me-review.json", home),
		"github.com/igor-kupczynski/help-me-review",
		[]string{"repo"},
	)
	if err != nil {
		log.Fatalf("review-helper: %v\n", err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: t.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	// list all the files in a pull request
	files, _, err := client.PullRequests.ListFiles(ctx, org, repo, pr, &github.ListOptions{PerPage: 1000})
	if err != nil {
		log.Fatalf("review-helper: %v\n", err)
	}

	sort.Sort(sort.Reverse(byChanges(files)))
	for _, v := range files {
		fmt.Printf("%v -> %v\n", *v.Filename, *v.Changes)
	}
}
