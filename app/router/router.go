package router

import (
	"fmt"
	"gin-learn-todo/app/controller"
	"gin-learn-todo/app/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	router := gin.New()

	router.RedirectFixedPath = true
	fmt.Println("base path:", router.BasePath())
	
	// 中间件
	router.Use(middleware.Login("admin"), gin.Recovery(), middleware.Cors())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "the router not exists",
		})
		return
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "the method not exists",
		})
		return
	})

	router.GET("ping", controller.GetPing)

	initRouter(router)

	return router
}

func initRouter(r *gin.Engine) {
	MyTest{}.Load(r)
	Api{}.load(r)
}
