package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/models"

type AchievementDAO interface {
	Create(m models.Achievement)
	Read(m models.Achievement)
	Update(m models.Achievement)
	Delete(m models.Achievement)
}
