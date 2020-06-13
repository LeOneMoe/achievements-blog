package models

import "time"

type User struct {
	Id        int64
	NickName  string
	FirstName string
	LastName  string
}

type Achievement struct {
	id              int64
	Title           string
	AchievementText string
	UTCdate         time.Time
	AuthorID        int64
}

type Like struct {
	id            int64
	userID        int64
	achievementID int64
}

type Comment struct {
	id            int64
	userID        int64
	achievementID int64
	commentText   string
}
