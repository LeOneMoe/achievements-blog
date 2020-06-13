package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/models"

type LikeDAO interface {
	Create(m models.Like)
	Read(m models.Like)
	Delete(m models.Like)
}
