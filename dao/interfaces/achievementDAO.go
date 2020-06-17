package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/models"

type AchievementDAO interface {
	Create(a *models.Achievement) error
	Update(id int64, a *models.Achievement) error
	Delete(id int64) error
	GetAll() ([]models.Achievement, error)
	GetByID(id int64) (models.Achievement, error)
}
