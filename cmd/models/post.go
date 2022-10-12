package models

import (
	"time"
)

type Post struct {
	ID     int64   `json:"id" validate:"omitempty"`
	ParentID uint `json:"parentId" gorm:"column:parent_id" validate:"omitempty"`
	UserID int64 `json:"userId" gorm:"column:user_id" validate:"omitempty"`
	Title string `json:"title"`
	Slug string `json:"slug" validate:"omitempty"`
	Summary string `json:"summary"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at" validate:"omitempty"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at" validate:"omitempty"`
	PublishedAt time.Time `json:"publishedAt" gorm:"column:published_at" validate:"omitempty"`
	Content string `json:"content"`
	IsPublished bool `json:"isPublished" gorm:"column:is_published"`
}



type Comment struct {
	ID     uint   `json:"id" pg:"id"`
	ParentID uint `json:"parentId" pg:"parent_id"`
	PostID uint `json:"postId" pg:"post_id"`
	Title string `json:"title" pg:"title"`
	Content string `json:"content" pg:"content"`
	PublishedAt time.Time `json:"publishedAt" pg:"published_at"`
}

type Tag struct {
	ID     uint   `json:"id" gorm:"primary_key;auto_increment;notNull"`
	Title string `json:"title" gorm:"notNull"`
	MetaTitle string `json:"metaTitle" gorm:"default:null"`
	Content string `json:"content"`
}

type PostTag struct {
	PostID Post `json:"postId" gorm:"notNull"`
	TagID Tag `json:"tagId" gorm:"notNull"`
}

type Category struct {
	ID     uint   `json:"id" gorm:"primary_key;auto_increment;notNull"`
	ParentID *Category `json:"parentId"`
	Title string `json:"title" gorm:"notNull"`
	MetaTitle string `json:"metaTitle" gorm:"notNull"`
	Slug string `json:"slug" gorm:"unique;notNull"`
	Content string `json:"content"`
}

type PostCategory struct {
	PostID Post `json:"postId" gorm:"notNull"`
	CategoryID Category `json:"categoryID" gorm:"notNull"`
}

type PostMeta struct {
	ID     uint   `json:"id" gorm:"primary_key;auto_increment;notNull"`
	PostID Post `json:"postId" gorm:"notNull"`
	Key string `json:"key" gorm:"notNull"`
	Content string `json:"content"`
}