package engine

import (
	"fmt"
	"log"
	"mini_spider/core/fetcher"
)

type workEngine struct{}

func (w workEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range requests {
			fmt.Println(item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching url: %v", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error"+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err

	}
	return r.ParserFunc(body), nil
}
