package model

type Comment struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	UserID         uint   `json:"user_id"`
	IllustrationID uint   `json:"illustration_id"`
	Content        string `json:"content"`
}

type CommentResponse struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	UserID         uint   `json:"user_id"`
	IllustrationID uint   `json:"illustration_id"`
	Content        string `json:"content"`
}