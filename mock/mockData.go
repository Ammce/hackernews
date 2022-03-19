package mocked_data

import "github.com/Ammce/hackernews/graph/models"

var MockUser = models.User{
	ID:       "123",
	Username: "Amel Muminovic",
	Email:    "amcenp@gmail.com",
}

var CommentMock1 = models.Comment{
	ID:          "224",
	NewsId:      "345",
	Text:        "This is serious",
	CreatedById: "123",
	CreatedAt:   "21/10/2022",
}

var CommentMock2 = models.Comment{
	ID:          "224",
	NewsId:      "345",
	Text:        "What the hell",
	CreatedById: "123",
	CreatedAt:   "21/10/2022",
}

var AllComments = []*models.Comment{&CommentMock1, &CommentMock2}

var MockNews1 = models.News{
	ID:           "345",
	Title:        "Russia attacks Ukraine",
	Text:         "Russia is about to attack Kyiv and all other major cities in Ukraine.",
	CreatedById:  "123",
	CreatedAt:    "20/10/2022",
	Published:    true,
	ApprovedById: "123",
}

var MockNews2 = models.News{
	ID:           "555",
	Title:        "Dol;ar is rising",
	Text:         "Due to the war crimes, dollar is rising up.",
	CreatedById:  "123",
	CreatedAt:    "20/10/2022",
	Published:    false,
	ApprovedById: "123",
}

var News = []*models.News{&MockNews1, &MockNews2}
