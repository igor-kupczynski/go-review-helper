package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
  "sort"
  "os"
  "strconv"
)

type tokenSpec struct {
	Token string `json:"token"`
}

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
  org, repo := os.Args[1], os.Args[2]
  pr, err := strconv.Atoi(os.Args[3])
  if err != nil {
    panic(err)
  }

	content, err := ioutil.ReadFile("token.json")
	if err != nil {
		panic(err)
	}

	spec := &tokenSpec{}
	if err := json.Unmarshal(content, &spec); err != nil {
		panic(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: spec.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	// list all the files in a pull request
  files, _, err := client.PullRequests.ListFiles(org, repo, pr, &github.ListOptions{})
  if err != nil {
    panic(err)
  }

  sort.Sort(sort.Reverse(byChanges(files)))
  for _, v := range files {
    fmt.Printf("%v -> %v\n", *v.Filename, *v.Changes)
  }
}
