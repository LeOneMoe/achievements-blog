package model

import "time"

type User struct {
	id        int64
	nickName  string
	firstName string
	lastName  string
}

type Achievement struct {
	id              int64
	title           string
	achievementText string
	UTCdate         time.Time
	authorID        int64
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
