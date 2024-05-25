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

	r.Any("/konnichiwa", func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
	})

	r.GET("/kyopro", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "https://atcoder.jp/users/brqnko",
		})
	})

	// 演習問題はここに追記してください

	r.Run(":8080") // ポート8080でサーバーを起動
}
