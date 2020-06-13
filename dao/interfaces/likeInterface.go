package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/model"

type LikeDAO interface {
	Create(m model.Like)
	Read(m model.Like)
	Delete(m model.Like)
}
