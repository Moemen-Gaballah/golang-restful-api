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

	g.JSON(http.StatusOK, gin.H{
		"message": "All Posts",
		"data":    posts,
	})
}

func Show(g *gin.Context) {
	post := getById(g)
	if post.ID == 0 {
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"message": "",
		"data":    post,
	})

}

func Store(g *gin.Context) {
	var post Post
	if err := g.ShouldBindJSON(&post); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": ""})
		return
	}
	post.Status = "Active"
	db.Create(&post)

	g.JSON(http.StatusCreated, gin.H{
		"message": "Post has been created",
		"data":    post,
	})
}

func Update(g *gin.Context) {
	oldPost := getById(g)
	if oldPost.ID == 0 {
		return
	}

	var requestPost Post
	if err := g.ShouldBindJSON(&requestPost); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": ""})
		return
	}

	oldPost.Title = requestPost.Title
	oldPost.Description = requestPost.Description
	if requestPost.Status != "" {
		oldPost.Status = requestPost.Status
	}

	db.Save(&oldPost)

	g.JSON(http.StatusOK, gin.H{
		"message": "Post has been updated",
		"data":    oldPost,
	})
}

func Delete(g *gin.Context) {
	post := getById(g)
	if post.ID == 0 {
		return
	}
	db.Delete(&post) // for hard delete db.Unscoped().Delete(&post)
	g.JSON(http.StatusOK, gin.H{
		"message": "Post has been deleted",
		"data":    "",
	})
}

func getById(g *gin.Context) Post {
	id := g.Param("id")
	var post Post
	db.First(&post, id)

	if post.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "We not found this post",
			"data":    "",
		})
	}
	return post
}
