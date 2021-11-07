package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 由c来创建响应
// H为字典

func sayHello(c *gin.Context) {
	// 返回json格式数据
	c.JSON(200, gin.H{
		"message": "Hello golang!",
	})

}

func main() {
	// 创建一个新的默认的路由引擎
	r := gin.Default()

	// 用户使用get访问/hello时，执行sayHello函数
	r.GET("/hello", sayHello)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	//启动服务
	r.Run(":9090")

}
