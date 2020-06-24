package models

type Achievement struct {
	ID              int64
	Title           string
	AchievementText string
	UTCDate         int64
	AuthorID        int64
}
