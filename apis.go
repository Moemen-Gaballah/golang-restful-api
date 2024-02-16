package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Posts(g *gin.Context) {
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(g.DefaultQuery("offset", "0"))

	var posts []Post
	db.Limit(limit).Offset(offset).Find(&posts)

	g.JSON(http.StatusCreated, gin.H{
		"error": "",
		"data":  posts,
	})
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
