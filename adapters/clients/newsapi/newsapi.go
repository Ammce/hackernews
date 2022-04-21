package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	externalArticle "github.com/Ammce/hackernews/domain/externalArticle"
)

type NewsApi struct{}

type Response struct {
	Status       string                            `json:"status"`
	TotalResults int16                             `json:"totalResults"`
	Articles     []externalArticle.ExternalArticle `json:"articles"`
}

func (n NewsApi) GetExternalArticlesByTopics(topics []string) ([]*externalArticle.ExternalArticlesByTopic, error) {
	var externalArticlesByTopic []*externalArticle.ExternalArticlesByTopic

	eaChan := make(chan externalArticle.ExternalArticlesByTopic, len(topics))

	for _, topic := range topics {
		go getArticlesByTopic(topic, eaChan)
	}

	// TODO - work this out using sync group
	data := <-eaChan
	data1 := <-eaChan
	externalArticlesByTopic = append(externalArticlesByTopic, &data, &data1)

	// for externalArticleByTopic := range eaChan {
	// 	externalArticlesByTopic = append(externalArticlesByTopic, &externalArticleByTopic)
	// }

	return externalArticlesByTopic, nil
}

func getArticlesByTopic(topic string, ch chan externalArticle.ExternalArticlesByTopic) {

	token := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", topic, token)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	var response = new(Response)
	sb := string(body)
	err3 := json.Unmarshal([]byte(sb), &response)
	if err3 != nil {
		fmt.Println("Error unmarshalling")
	}

	ch <- externalArticle.ExternalArticlesByTopic{
		Topic:    topic,
		Articles: response.Articles,
	}
}

func (n NewsApi) GetTopExternalArticlesPerCountry(country *string) ([]*externalArticle.ExternalArticle, error) {
	token := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", *country, token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	var response = new(Response)
	sb := string(body)
	err3 := json.Unmarshal([]byte(sb), &response)
	if err3 != nil {
		fmt.Println("Error unmarshalling")
	}
	var externalArticles []*externalArticle.ExternalArticle

	for _, asd := range response.Articles {
		newAr := asd
		externalArticles = append(externalArticles, &newAr)
	}

	return externalArticles, nil
}

func NewNewsApi() NewsApi {
	return NewsApi{}
}