package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"news-grabber-api/config"
)

type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type sourcesAPI struct {
	Sources []Source `json:"sources"`
}

type topicsAPI struct {
	Articles []Topic `json:"articles"`
}

func GetSources(category string) []Source {
	body := getData(sourceURL(category))

	var sourceApi sourcesAPI
	json.Unmarshal(body, &sourceApi)

	return sourceApi.Sources
}

func getTopics(sources []Source) []Topic {
	var topics []Topic

	for _, source := range sources {
		body := getData(articleURL(source.ID))

		var topicApi topicsAPI
		json.Unmarshal(body, &topicApi)
		topics = append(topics, topicApi.Articles...)
	}
	return topics
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return body
}

func sourceURL(category string) string {
	return fmt.Sprintf("https://newsapi.org/v2/sources?language=en&category=%s", category)
}

func articleURL(id string) string {
	return fmt.Sprintf("https://newsapi.org/v2/articles?source=%s&sortBy=latest&apiKey=%s", id, config.ApiKey)
}
