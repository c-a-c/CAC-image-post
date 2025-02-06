package model

type Like struct {
	ID             uint `json:"id" gorm:"primary_key"`
	UserID         uint `json:"user_id"`
	IllustrationID uint `json:"illustration_id"`
}

type LikeResponse struct {
	ID             uint `json:"id" gorm:"primary_key"`
	UserID         uint `json:"user_id"`
	IllustrationID uint `json:"illustration_id"`
}