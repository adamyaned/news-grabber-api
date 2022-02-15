package main

import (
	"news-grabber-api/news"
	"news-grabber-api/router"
)

func main() {
	r := router.New()
	a := news.New()

	go a.Serve()
	
	r.Run()
}
