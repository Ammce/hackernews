package mocked_data

import "github.com/Ammce/hackernews/models"

var MockUser = models.User{
	ID:    "123",
	Name:  "Amel Muminovic",
	Email: "amcenp@gmail.com",
}

var MockNews1 = models.News{
	ID:          "345",
	Title:       "Russia attacks Ukraine",
	Text:        "Russia is about to attack Kyiv and all other major cities in Ukraine.",
	CreatedById: "123",
	CreatedAt:   "20/10/2022",
	Published:   true,
}

var MockNews2 = models.News{
	ID:          "555",
	Title:       "Dol;ar is rising",
	Text:        "Due to the war crimes, dollar is rising up.",
	CreatedById: "123",
	CreatedAt:   "20/10/2022",
	Published:   false,
}

var News = []*models.News{&MockNews1, &MockNews2}
