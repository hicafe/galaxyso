package main

import (
	"service"
	"time"
)

func init() {
	service.InitSourceUrls() //init source urls to be crawl
}

/*
Spider
*/
func main() {
	for {
		service.Crawler()
		time.Sleep(time.Hour * 24)
	}
}
