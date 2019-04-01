package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PaginationForm :
type PaginationForm struct {
	Offset int `form:"offset,default=0" binding:"min=0"`
	Limit  int `form:"limit,default=10" binding:"min=0"`
}

// ModelView :
type ModelView struct {
}

// ResponseFromError :
func ResponseFromError(ctx *gin.Context, err error) {
	switch err {
	case gorm.ErrRecordNotFound:
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
	default:
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
	}
}

// ResponseFromData :
func ResponseFromData(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"message": "ok",
		"data":    data,
	})
}
