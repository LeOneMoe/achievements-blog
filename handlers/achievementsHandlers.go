package handlers

import "github.com/gin-gonic/gin"

func GetAllAchievements(c *gin.Context) {
	achievements, err := achievementDAO.GetAll()
	if err != nil {
		c.JSON(520, gin.H{
			"achievements": nil,
		})
	}

	c.JSON(200, gin.H{
		"achievements": achievements,
	})
}
