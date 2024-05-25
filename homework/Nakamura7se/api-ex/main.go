package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tweetの構造体を定義
type Tweet struct {
	ID   int    `json:"id"`   // JSONのidがTweetのIDにマッピング(対応)される
	Text string `json:"text"` // JSONのtextがTweetのTextにマッピング(対応)される
}

// Tweetのスライスを定義。ここの変数にツイートが格納される
var tweets = []Tweet{
	{ID: 1, Text: "Hello, World!"},
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
		newTweet.ID = nextID                 // IDを設定
		nextID++                             // 次のIDを+1
		tweets = append(tweets, newTweet)    // ツイートを追加
		c.JSON(http.StatusCreated, newTweet) // HTTPステータスコード201(Created)で、新しいツイートを返す
	})

	// Read機能以降はこの下に追記していきます

	r.GET("/tweets/:id", func(c *gin.Context) { // HTTP GETで/tweetsにアクセスしたときの処理
		c.JSON(http.StatusOK, tweets) // HTTPステータスコード200(OK)で、ツイート一覧を返す
	})

	r.Run(":8080") // ポート8080でサーバーを起動
}
