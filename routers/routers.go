package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/levnzzz/ginEssential/controller"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)

	return r
}