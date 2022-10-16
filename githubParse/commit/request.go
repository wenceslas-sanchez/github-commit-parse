package commit

import (
	"fmt"
	"githubParse/githubParse/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const BaseRepository = utils.BaseURL + "/repos"
const BufferSize = 20

func Information(client *http.Client, owner, repository string) (*[]*Commit, error) {
	var mu sync.Mutex // link to commits below
	var commits []*Commit
	var responses = make(chan *http.Response, BufferSize)
	var errors = make(chan error, BufferSize)

	go parseResponses(&responses, &errors, &commits, &mu)

	res, err := requestCommitPage(client, owner, repository, "1")
	if err != nil {
		return nil, err
	}
	responses <- res

	lastPageNum, ok := parseLinkLastPage(&res.Header)
	if ok {
		fmt.Println(lastPageNum)
	}
	max, _ := strconv.Atoi(lastPageNum)
	for _, i := range makeRange(2, max+1) {
		res, err := requestCommitPage(client, owner, repository, strconv.Itoa(i))
		if err != nil {
			errors <- err
		}
		responses <- res
	}

	return &commits, nil
}

func requestCommitPage(client *http.Client, owner string, repository string, page string) (*http.Response, error) {
	log.Printf("request page %v", page)
	baseCommitsUrl := BaseRepository + "/" + owner + "/" + repository + "/commits?page=" + page
	res, err := utils.GetRequest(client, baseCommitsUrl)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func parseResponses(responses *chan *http.Response, errors *chan error, commits *[]*Commit, mu *sync.Mutex) {
	for {
		select {
		case res := <-*responses:
			var commit []*Commit
			parseResponse(res, &commit)
			mu.Lock()
			*commits = append(*commits, commit...)
			mu.Unlock()
		case err := <-*errors:
			log.Fatal(err)
		}
	}
}

func parseLinkLastPage(header *http.Header) (string, bool) {
	link := header.Get("Link")
	if link == "" {
		return "", false
	}
	lastUrl := strings.Replace(strings.Replace(strings.Split(link, ";")[1], "<", "", 1), ">", "", 1)
	lastPageNum := strings.Split(lastUrl, "page=")[1]

	return lastPageNum, true
}

func parseResponse(response *http.Response, commits *[]*Commit) {
	err := utils.DecodeResponseJSON(response, commits)
	if err != nil {
		fmt.Println(err)
	}
}

// generate a slice of all integer between min (included) and max (excluded)
func makeRange(min, max int) []int {
	nums := make([]int, max-min)
	for i := range nums {
		nums[i] = min + i
	}

	return nums
}

// get single commit reference details
func GetDetails(client *http.Client, owner, repository, ref string) (*Commit, error) {
	res, err := utils.GetRequest(client, BaseRepository+"/"+owner+"/"+repository+"/commits"+ref)
	if err != nil {
		return nil, err
	}

	var commit *Commit
	err = utils.DecodeResponseJSON(res, commit)
	if err != nil {
		return nil, fmt.Errorf("commit decode issue: %s", err)
	}

	return commit, nil
}
