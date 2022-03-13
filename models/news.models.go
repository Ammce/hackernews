package models

type News struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	Published   bool   `json:"published"`
	CreatedById string `json:"createdById"`
	CreatedAt   string `json:"createdAt"`
}
