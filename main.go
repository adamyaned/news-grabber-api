package main

import "news-grabber-api/router"

func main() {
	r := router.New()

	r.Run()
}
