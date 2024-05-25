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
	r.POST("/hello", func(c *gin.Context) { // HTTP POSTで/helloにアクセスしたときの処理
		c.JSON(http.StatusOK, gin.H{ // HTTPステータスコード200(OK)で、JSONを返す
			"message": "Hello, Zli!",
		})
	})

	r.DELETE("/bye", func(c *gin.Context) { // HTTP DELETEで/byeにアクセスしたときの処理
		c.JSON(http.StatusOK, gin.H{ // HTTPステータスコード200(OK)で、JSONを返す
			"message": "Bye!",
		})
	})

	r.GET("/konnichiwa", func(c *gin.Context) { // HTTP GETで/konnichiwaにアクセスしたときの処理
		c.JSON(http.StatusBadRequest, gin.H{ // HTTPステータスコード400(Bad Request)で、JSONを返す
			"message": "Bad Request",
		})
	})

	r.GET("/blender", func(c *gin.Context) { // HTTP GETで/pingにアクセスしたときの処理
		c.JSON(http.StatusOK, gin.H{ // HTTPステータスコード200(OK)で、JSONを返す
			"message": "Blender is exciting!",
		})
	})

	r.Run(":8080") // ポート8080でサーバーを起動
	//C-cでサーバーを停止
}
