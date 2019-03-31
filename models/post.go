package models

// Post model
type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Subtitle  string `json:"sub_title"`
	Slug      string `json:"slug"`
	Content   string `json:"content"`
	Isdeleted string `json:"is_deleted"`
	// Author    string `json:author`
}
