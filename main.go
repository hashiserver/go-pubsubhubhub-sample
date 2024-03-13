package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	r.GET("/feed", func(c *gin.Context) {
		// JSONレスポンスを返す
		hubChallenge := c.Query("hub.challenge")
		log.Println("hub.challenge", hubChallenge)
		c.String(200, hubChallenge)
	})

	r.POST("/feed", func(c *gin.Context) {
		log.Println("POST")
		log.Println(c.Request.Body)
		c.Status(http.StatusNoContent)
	})

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
