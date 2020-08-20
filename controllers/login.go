package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"github.com/hariNEzuMI928/run-together-towards-goals/slack"
)

// SessionInfo ...
type SessionInfo struct {
	UserID interface{}
}

// LoginInfo ...
var LoginInfo SessionInfo

// Login ...
func Login(c *gin.Context) {
	user, err := LoginUser(c)
	if err != "" {
		c.HTML(http.StatusFound, "login.html", gin.H{
			"err": err,
		})
		return
	}

	//セッションにデータを格納する
	session := sessions.Default(c)
	session.Set("UserID", int(user.ID))
	session.Set("UserNickname", user.Nickname)
	session.Save()

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"user": user,
	})
}

// Logout ...
func Logout(c *gin.Context) {
	//セッションからデータを破棄する
	session := sessions.Default(c)

	session.Clear()
	session.Save()

	log.Println("ログアウト")

	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "Bye!",
	})
}

// SessionCheck ...
func SessionCheck(c *gin.Context) {
	session := sessions.Default(c)
	LoginInfo.UserID = session.Get("UserID")

	if strings.HasPrefix(c.Request.RequestURI, "/api") {
		c.Next()
		return
	} else {
		if LoginInfo.UserID == nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"err": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// GetLoginUser ...
func GetLoginUser(c *gin.Context) models.User {
	r := models.NewUserRepository()
	user := r.GetLoginUser(sessions.Default(c).Get("UserID"))
	if user.ID == 0 {
		log.Println("cannot get GetLoginUser")
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"err": "cannot get GetLoginUser",
		})
		c.Abort()
	// c.Redirect(http.StatusMovedPermanently, "/logout")
	}

	return user
}

// GetLoginUserID ...
func GetLoginUserID(c *gin.Context) int {
	ru := models.NewUserRepository()
	user := ru.GetLoginUser(sessions.Default(c).Get("UserID"))
	userID := int(user.ID)
	if userID == 0 {
		log.Println("cannot get GetLoginUserID")
		// c.Redirect(http.StatusMovedPermanently, "/logout")
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"err": "cannot get GetLoginUserID",
		})
		c.Abort()
	}

	return userID
}

// ForgotPassword ...
func ForgotPassword(c *gin.Context) {
	nickname, _ := c.GetPostForm("nickname")
	slack.NoticeForgotPass(nickname)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "Plz wait for setting new password",
	})

}
