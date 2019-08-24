package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/service"
)

type GinServer struct {
	tweetManager *service.TweetManager
}

type GinTweet struct {
	User          string `json:"user"`
	Text          string `json:"text"`
	QuotedTweetId int    `json:"quotedTweetId"`
}

func NewGinServer(tweetManager *service.TweetManager) *GinServer {
	return &GinServer{tweetManager: tweetManager}
}

func (ginServer *GinServer) StartServer() {
	router := gin.Default()

	router.POST("/publishTextTweet", ginServer.publishTextTweet)

	router.POST("/publishQuotedTweet", ginServer.publishQuotedTweet)

	router.Run()
}

func (server *GinServer) publishTextTweet(c *gin.Context) {

	var tweet GinTweet

	if errBind := c.ShouldBindJSON(&tweet); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	textTweet := domain.NewTweet(tweet.User, tweet.Text)

	server.tweetManager.PublishTweet(textTweet)

	id, errPublish := server.tweetManager.PublishTweet(textTweet)

	if errPublish != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errPublish.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (server *GinServer) publishQuotedTweet(c *gin.Context) {

	var tweet GinTweet

	if errBind := c.ShouldBindJSON(&tweet); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	quotedTweet := server.tweetManager.GetTweetById(tweet.QuotedTweetId)

	quoteTweet := domain.NewQuotedTweet(tweet.User, tweet.Text, quotedTweet)

	//server.tweetManager.PublishTweet(quoteTweet)

	id, errPublish := server.tweetManager.PublishTweet(quoteTweet)

	if errPublish != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errPublish.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
