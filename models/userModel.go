package models

import "time"

type User struct {
	ID        int64
	NickName  string
	FirstName string
	LastName  string
}

type Achievement struct {
	ID              int64
	Title           string
	AchievementText string
	UTCdate         time.Time
	AuthorID        int64
}

type Like struct {
	ID            int64
	UserID        int64
	AchievementID int64
}

type Comment struct {
	ID            int64
	UserID        int64
	AchievementID int64
	CommentText   string
}
