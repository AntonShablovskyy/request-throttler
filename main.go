package main

import (
	"flag"
	"request-throttler/internal/throttler"
)

func main() {
	maxQPS, requests, url := parseArgs()

	t := throttler.New(maxQPS, requests)
	t.Start(url)
}

func parseArgs() (maxQPS, requests int, url string) {
	flag.IntVar(&maxQPS, "maxQPS", 100, "Maximum queries per second")
	flag.IntVar(&requests, "requestsCount", 1000, "Total requests count")
	flag.StringVar(&url, "url", "http://localhost:8080", "Url to throttle")

	flag.Parse()

	return maxQPS, requests, url
}
