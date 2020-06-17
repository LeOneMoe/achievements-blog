package handlers

import "github.com/LeOneMoe/go-gin-react-crud/dao/factories"
import "github.com/LeOneMoe/go-gin-react-crud/utilities"

var cfg, _ = utilities.GetConfiguration()

var (
	userDAO        = factories.UserFactory(cfg.Driver)
	achievementDAO = factories.AchievementFactory(cfg.Driver)
	commentDAO     = factories.CommentFactory(cfg.Driver)
	likerDAO       = factories.LikeFactory(cfg.Driver)
)
