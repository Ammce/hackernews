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
	//TODO - Call len(topics) amount of Go routines to access the data from different media. Use channels
	return nil, nil
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
