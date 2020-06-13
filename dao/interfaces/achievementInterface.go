package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/model"

type AchievementDAO interface {
	Create(m model.Achievement)
	Read(m model.Achievement)
	Update(m model.Achievement)
	Delete(m model.Achievement)
}
