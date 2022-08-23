package users

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/MyApp")
	routes.POST("/DisplayUser", h.DisplayUser)
	routes.GET("/DisplayAllUsers", h.DisplayAllUsers)
}
