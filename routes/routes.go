package routes

import (
	"github.com/daisuzuki829/run_together_towards_goals/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Handler(dbConn *gorm.DB) {

	handler := controllers.Handler{
		Db: dbConn,
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")

	r.GET("/users", handler.GetAllUsers)
	rUser := r.Group("/user")
	{
		rUser.POST("/add", handler.AddUser)
		rUser.GET("/edit/:id", handler.GetUser)
		rUser.POST("/edit_ok/:id", handler.EditUser)
		rUser.GET("/delete/:id", handler.DeleteUser)
		rUser.DELETE("/delete/:id", handler.DeleteUser)
	}

	r.GET("/genre", handler.GetAllGenres)
	rGenre := r.Group("/genre")
	{
		rGenre.POST("/add", handler.AddGenre)
		rGenre.GET("/edit/:id", handler.GetGenre)
		rGenre.POST("/edit_ok/:id", handler.EditGenre)
		rGenre.GET("/delete/:id", handler.DeleteGenre)
		rGenre.DELETE("/delete/:id", handler.DeleteGenre)
	}

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"title": "TITLE",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	//spew.Dump(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
