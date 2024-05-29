package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tweet struct {
	ID    int    `json:"id"`   // JSONのidがTweetのIDにマッピング(対応)される
	Text  string `json:"text"` // JSONのtextがTweetのTextにマッピング(対応)される
	Likes int    `json:"likes"`
}

// Tweetのスライスを定義。ここの変数にツイートが格納される
var tweets = []Tweet{
	{ID: 1, Text: "Hello, World!", Likes: 0},
	{ID: 2, Text: "Intro Go", Likes: 0},
}

var nextID = 3

func main() {
	r := gin.Default() // ginのルーターを作成

	r.GET("/hello", func(c *gin.Context) { // HTTP GETで/helloにアクセスしたときの処理
		c.JSON(http.StatusOK, gin.H{ // HTTPステータスコード200(OK)で、JSONを返す
			"message": "Hello, Zli!",
		})
	})
	// 演習問題はここに追記してください

	r.DELETE("/bye", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bye!",
		})
	})

	r.PUT("/konnichiwa", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad Request",
		})
	})

	r.GET("/tweets/:id", func(c *gin.Context) { // HTTP PUTで/tweets/:idにアクセスしたときの処理
		id, err := strconv.Atoi(c.Param("id")) // URLの:idを取得
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
		}

		for _, tweet := range tweets {
			if tweet.ID == id {
				c.JSON(http.StatusOK, tweet) // HTTPステータスコード200(OK)で、ツイートを返す
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"}) // HTTPステータスコード404(Not Found)で、ツイートが見つからないエラーメッセージを返す
	})

	r.GET("/tweets/:id/like", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id")) // URLの:idを取得
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
		}

		/*for _, tweet := range tweets {
			if tweet.ID == id {
				tweet.Likes++
				c.JSON(http.StatusOK, tweet)
				return
			}
		}*/

		for i, tweet := range tweets {
			if tweet.ID == id {
				tweets[i].Likes++
				c.JSON(http.StatusOK, tweets[i])
				return
			}
		}

	})

	r.Run(":8080") // ポート8080でサーバーを起動
}
