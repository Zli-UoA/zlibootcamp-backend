package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

// Tweetの構造体を定義
type Tweet struct {
	ID   int    `json:"id"`   // JSONのidがTweetのIDにマッピング(対応)される
	Text string `json:"text"` // JSONのtextがTweetのTextにマッピング(対応)される
	Like int    `json:"like"` // JSONのidがTweetのIDにマッピング(対応)される
}

// Tweetのスライスを定義。ここの変数にツイートが格納される
var tweets = []Tweet{
	{ID: 1, Text: "Hello, World!", Like: 0},
}

var nextID = 2

func main() {
	r := gin.Default() // ginのルーターを作成

	r.POST("/tweets", func(c *gin.Context) { // HTTP POSTで/tweetsにアクセスしたときの処理
		var newTweet Tweet
		if err := c.ShouldBindJSON(&newTweet); err != nil { // リクエストボディをTweet構造体にバインド。エラーがあればエラーメッセージを返す
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTweet.Like = 0
		newTweet.ID = nextID                 // IDを設定
		nextID++                             // 次のIDを+1
		tweets = append(tweets, newTweet)    // ツイートを追加
		c.JSON(http.StatusCreated, newTweet) // HTTPステータスコード201(Created)で、新しいツイートを返す
	})

	r.GET("/tweets", func(c *gin.Context) { // HTTP GETで/tweetsにアクセスしたときの処理
		c.JSON(http.StatusOK, tweets) // HTTPステータスコード200(OK)で、ツイート一覧を返す
	})

	r.PUT("/tweets/:id", func(c *gin.Context) { // HTTP PUTで/tweets/:idにアクセスしたときの処理
		id, err := strconv.Atoi(c.Param("id")) // URLの:idを取得
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
		}
		for i, tweet := range tweets {
			if tweet.ID == id {
				var updatedTweet Tweet
				if err := c.ShouldBindJSON(&updatedTweet); err != nil { // リクエストボディをTweet構造体にバインド。エラーがあればエラーメッセージを返す
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				updatedTweet.ID = id
				tweets[i] = updatedTweet            // ツイートを更新
				c.JSON(http.StatusOK, updatedTweet) // HTTPステータスコード200(OK)で、更新されたツイートを返す
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"}) // HTTPステータスコード404(Not Found)で、ツイートが見つからないエラーメッセージを返す
	})

	r.DELETE("/tweets/:id", func(c *gin.Context) { // HTTP DELETEで/tweets/:idにアクセスしたときの処理
		id, err := strconv.Atoi(c.Param("id")) // URLの:idを取得
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
		}
		for i, tweet := range tweets {
			if tweet.ID == id {
				tweets = append(tweets[:i], tweets[i+1:]...)             // ツイートを削除
				c.JSON(http.StatusOK, gin.H{"message": "Tweet deleted"}) // HTTPステータスコード200(OK)で、ツイートが削除されたメッセージを返す
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"}) // HTTPステータスコード404(Not Found)で、ツイートが見つからないエラーメッセージを返す
	})

	r.GET("/tweets/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) // URLの:idを取得
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
		}

		for _, tweet := range tweets {
			if tweet.ID == id { // ツイートを更新
				c.JSON(http.StatusOK, tweet) // HTTPステータスコード200(OK)で、更新されたツイートを返す
				return
			}
		}
	})

	r.POST("/tweets/:id/like", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
		}

		for i, tweet := range tweets {
			if tweet.ID == id {
				tweets[i].Like++                   // ツイートを更新
				ctx.JSON(http.StatusOK, tweets[i]) // HTTPステータスコード200(OK)で、更新されたツイートを返す
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
	})

	r.Run(":8080") // ポート8080でサーバーを起動
}
