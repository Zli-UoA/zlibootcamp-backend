package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //ginのルーターを作成

	r.GET("/hello", func(c *gin.Context) { //http.getで/helloにアクセスしたときの処理
		c.JSON(http.StatusOK, gin.H{ //HTTPステータスコード200でjsonを返す
			"message": "Hello,World!",
		})
	})

	r.Run(":8080") //ポート8080でサーバー起動
}
