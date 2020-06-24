package models

type Comment struct {
	ID            int64
	UserID        int64
	AchievementID int64
	CommentText   string
}
