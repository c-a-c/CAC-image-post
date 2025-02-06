package model

import "time"

type Illustration struct {
	ID		  uint       `json:"id" gorm:"primary_key"`
	ImageURL  string     `json:"image_url"`
	Title     string     `json:"title"`
	User      User       `json:"user"`
	UserID    uint       `json:"user_id"`
	Caption   string     `json:"caption"`
	AgeRestriction  bool `json:"age_restriction"`
	Tags	  []Tag      `json:"tags" gorm:"many2many:illustration_tags;"`
	TimeStamp time.Time  `json:"time_stamp"`
	Likes	  []Like     `json:"likes" gorm:"many2many:illustration_likes;"`
}

type IllustrationResponse struct {
	ID		  uint       `json:"id" gorm:"primary_key"`
	ImageURL  string     `json:"image_url"`
	Title     string     `json:"title"`
	User      User       `json:"user"`
	UserID    uint       `json:"user_id"`
	Caption   string     `json:"caption"`
	AgeRestriction  bool `json:"age_restriction"`
	Tags	  []Tag      `json:"tags" gorm:"many2many:illustration_tags;"`
	TimeStamp time.Time  `json:"time_stamp"`
	Likes	  []Like     `json:"likes" gorm:"many2many:illustration_likes;"`
}