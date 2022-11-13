package models

type Todo struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	User   User   `json:"user"`
}
