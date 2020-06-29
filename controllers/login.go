package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type SessionInfo struct {
	UserId interface{}
}

var LoginInfo SessionInfo

// Login...
func Login(c *gin.Context) {
	user, err := LoginUser(c)

	if err != "" {
		c.HTML(http.StatusFound, "login.html", gin.H{
			"err": err,
		})
		return
	}

	UserId := user.ID
	UserNickname := user.Nickname

	//セッションにデータを格納する
	session := sessions.Default(c)
	session.Set("UserId", UserId)
	session.Set("UserNickname", UserNickname)
	session.Save()

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"user": user,
	})
}

// Logout...
func Logout(c *gin.Context) {
	//セッションからデータを破棄する
	session := sessions.Default(c)

	session.Clear()
	session.Save()
	LoginInfo.UserId = session.Get("UserId")

	log.Println("ログアウト")

	c.HTML(http.StatusOK, "login.html", gin.H{
		"logout": "Bye!",
	})
}

// SessionCheck...
func SessionCheck(c *gin.Context) {
	session := sessions.Default(c)
	LoginInfo.UserId = session.Get("UserId")

	if strings.HasPrefix(c.Request.RequestURI, "/api") {
		c.Next()
		return
	} else {
		if LoginInfo.UserId == nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"err": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
