package routes

import (
	"github.com/clementejuliana/grupoDprojetoGO/server/controller"
	"github.com/gin-gonic/gin"
)

func Config(router *gin.Engine) *gin.Engine{
	main := router.Group("api/v1")
	{
	  produtos := main.Group("produtos")
	  {
		produtos.GET("/", controllers.produtoShow)
	}
}
	
return router

}