package users

import (
	"fmt"
	"net/http"
	"sirka-be-test/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type DisplayUserRequestBody struct {
	Userid	string `json:"userid"`
}

func (h handler) DisplayUser(c *gin.Context) {
	body := DisplayUserRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	fmt.Println(body)

	id := body.Userid

	var user models.Users

	if result := h.DB.Where("Userid = ?", id).First(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &user)
}
