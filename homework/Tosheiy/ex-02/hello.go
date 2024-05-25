package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
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
		c.JSON(http.StatusBadRequest, gin.H{ // HTTPステータスコード400(Bad)で、JSONを返す
			"error": "Bad Request",
		})
	})

	r.GET("/konnichiwa/:code", func(c *gin.Context) { // HTTP GETで/konnichiwaにアクセスしたときの処理
		code, err := strconv.Atoi(c.Param("code"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"eror": "Bad Request"})
			return
		}

		if code == 123 {
			c.JSON(http.StatusOK, gin.H{"password": 322})
			return
		} else if code == 322 {
			c.JSON(http.StatusOK, gin.H{"status": "clear"})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "code not found."})
	})

	r.Run(":8080") // ポート8080でサーバーを起動
}
