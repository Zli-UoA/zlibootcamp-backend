package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
)

func main() {
	r := gin.Default() // ginのルーターを作成

	r.GET("/hello", func(c *gin.Context) { // HTTP GETで/helloにアクセスしたときの処理
		c.JSON(http.StatusOK, gin.H{ // HTTPステータスコード200(OK)で、JSONを返す
			"message": "Hello, Zli!",
		})
	})

	// 演習問題はここに追記してください

	r.DELETE("/bye", func(c *gin.Context) { // HTTP DELETEで/byeにアクセスしたときの処理
		c.JSON(http.StatusOK, gin.H{ // HTTPステータスコード200(OK)で、JSONを返す
			"message": "Bye!",
		})
	})

	r.GET("/konnichiwa", func(c *gin.Context) { // HTTP GETで/konnichiwaにアクセスしたときの処理
		// fmt.Println(gin.Context.Accepted);
		c.JSON(http.StatusBadRequest, gin.H{ // HTTPステータスコード400(Bad)で、JSONを返す
			"message": "Bad Request",
		})
	})

	r.Run(":8080") // ポート8080でサーバーを起動
}
