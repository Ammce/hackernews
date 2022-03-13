package models

type Comment struct {
	ID          string `json:"id"`
	Text        string `json:"name"`
	CreatedById string `json:"createdById"`
	NewsId      string `json:"newsId"`
	CreatedAt   string `json:"createdAt"`
}
