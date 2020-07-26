package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

// GetGenres...
func (h *Handler) GetAllGenres(c *gin.Context) {
	r := models.NewGenreRepository()
	genres := r.GetAll()

	c.HTML(http.StatusOK, "genres.html", gin.H{
		"genres": genres,
	})
}

// AddGenres...
func (h *Handler) AddGenre(c *gin.Context) {
	r := models.NewGenreRepository()
	r.GenreName, _ = c.GetPostForm("genre_name")
	r.Add(&r)

	c.Redirect(http.StatusMovedPermanently, "/_genres")
}

// GetGenres...
func (h *Handler) GetGenre(c *gin.Context) {
	r := models.NewGenreRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	genre := r.GetOne(id)
	c.HTML(http.StatusOK, "genre_edit.html", gin.H{
		"genre": genre,
	})
}

// EditGenres...
func (h *Handler) EditGenre(c *gin.Context) {
	r := models.NewGenreRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	genre := r.GetOne(id)
	genre.GenreName, _ = c.GetPostForm("genreName")
	r.Edit(genre)
	c.Redirect(http.StatusMovedPermanently, "/_genres")
}

// DeleteGenres...
func (h *Handler) DeleteGenre(c *gin.Context) {
	r := models.NewGenreRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/_genres")
}
