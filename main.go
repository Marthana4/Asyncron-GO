package main

import (
	"github.com/Marthana4/web5sec/services"
	"time"
)

func main() {
	titleConfig := services.Config{
		URL:           "https://go.dev/doc/",
		FetchInterval: 5 * time.Second,
	}

	dataConfig := services.Config{
		URL:           "https://go.dev/doc/",
		FetchInterval: 3 * time.Second,
	}

	go services.FetchTitle(titleConfig)
	go services.FetchData(dataConfig)

	select {}
}
