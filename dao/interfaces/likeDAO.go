package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/models"

type LikeDAO interface {
	Create(l *models.Like) error
	Delete(achID, userID int64) error
	GetCount(achID int64) (int, error)
}
