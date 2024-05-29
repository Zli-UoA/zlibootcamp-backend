package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // ginのルーターを作成

	r.GET("/hello", func(c *gin.Context) { // HTTP GETで/helloにアクセスしたときの処理
		c.JSON(http.StatusOK, gin.H{ // HTTPステータスコード200(OK)で、JSONを返す
			"message": "Hello, World!",
		})
	})

	// 演習問題はここに追記してください
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Zli!",
		})
	})

	r.DELETE("/bye", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bye!",
		})
	})

	r.GET("/konnichiwa", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	})
}
