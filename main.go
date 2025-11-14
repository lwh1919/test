package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置 Gin 为 release 模式（生产环境）
	gin.SetMode(gin.ReleaseMode)

	// 创建 Gin 路由引擎
	r := gin.Default()

	// Ping-pong 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// 根路径
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin!",
			"endpoints": []string{
				"/ping",
				"/health",
			},
		})
	})

	// 启动服务器，监听 8000 端口
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
