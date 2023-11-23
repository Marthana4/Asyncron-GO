package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"time"
)

type Config struct {
	URL           string
	FetchInterval time.Duration
}

func FetchTitle(config Config) {
	for {
		doc, err := goquery.NewDocument(config.URL)
		if err != nil {
			fmt.Printf("Error fetching title from %s: %v\n", config.URL, err)
			time.Sleep(config.FetchInterval)
			continue
		}

		title := doc.Find("title").Text()
		fmt.Printf("Title from %s: %s\n", config.URL,time.Now().Format("05"), title)

		time.Sleep(config.FetchInterval)
	}
}

func FetchData(config Config) {
	tick := time.Tick(config.FetchInterval)

	for range tick {
		doc, err := goquery.NewDocument(config.URL)
		if err != nil {
			fmt.Printf("Error fetching data from %s: %v\n", err)
			continue
		}

		data := doc.Find("h1").Text()
		fmt.Printf("Data from %s: %s\n", config.URL, time.Now().Format("05"), data)
	}
}
