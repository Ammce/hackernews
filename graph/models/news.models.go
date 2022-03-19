package models

type News struct {
	ID           string `json:"id"`
	Text         string `json:"text"`
	Title        string `json:"title"`
	Published    bool   `json:"published"`
	CreatedById  string `json:"createdById"`
	ApprovedById string `json:"approvedById"`
	CreatedAt    string `json:"createdAt"`
}
