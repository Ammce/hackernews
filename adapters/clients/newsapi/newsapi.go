package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	externalArticle "github.com/Ammce/hackernews/domain/externalArticle"
)

type NewsApi struct{}

type Response struct {
	Status       string                            `json:"status"`
	TotalResults int16                             `json:"totalResults"`
	Articles     []externalArticle.ExternalArticle `json:"articles"`
}

func (n NewsApi) GetTopArticlesPerCountry(country *string) ([]*externalArticle.ExternalArticle, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=4bd0597f548c4f279483863901d721c3", *country)
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
	var toMap []*externalArticle.ExternalArticle

	for _, asd := range response.Articles {
		newAr := asd
		toMap = append(toMap, &newAr)
	}

	return toMap, nil
}

func NewNewsApi() NewsApi {
	return NewsApi{}
}
