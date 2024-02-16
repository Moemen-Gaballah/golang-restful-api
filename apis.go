package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Posts(g *gin.Context) {
	// TODO
}

func Show(g *gin.Context) {

}

func Store(g *gin.Context) {

	var post Post
	if err := g.ShouldBindJSON(&post); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": ""})
		return
	}
	post.Status = "Active"
	db.Create(&post)

	g.JSON(http.StatusCreated, gin.H{
		"error": "",
		"data":  post,
	})
}

func Update(g *gin.Context) {

}

func Delete(g *gin.Context) {

}
