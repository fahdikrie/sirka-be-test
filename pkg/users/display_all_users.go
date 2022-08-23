package users

import (
	"net/http"
	"sirka-be-test/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DisplayAllUsers(c *gin.Context) {
	var users []models.Users

	if 	result := h.DB.Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &users)
}
