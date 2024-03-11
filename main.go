package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	// ルートURL ("/") に対するGETリクエストをハンドル
	r.GET("/feed", func(c *gin.Context) {
		// JSONレスポンスを返す
		hubChallenge := c.Query("hub.challenge")
		c.String(200, hubChallenge)
	})

	r.POST("/feed", func(c *gin.Context) {
		log.Println("POST")
		log.Println(c.Request.Body)
	})

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
