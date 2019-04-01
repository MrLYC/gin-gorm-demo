package demo

import (
	"github.com/gin-gonic/gin"
)

func route(enging *gin.Engine) {
	enging.GET("/", IndexView)
	enging.GET("/pairs/:id", RetrievePairView)
	enging.DELETE("/pairs/:id", DeletePairView)
	enging.GET("/pairs/", ListPairView)
	enging.PUT("/pairs/", CreatePairView)
	enging.POST("/pairs/", UpdatePairView)
}
