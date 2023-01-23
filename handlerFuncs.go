package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func CreateWord(cc *mongo.Collection) gin.HandlerFunc {

	fn := func(ctx *gin.Context) {

		var word Word

		if err := ctx.ShouldBindJSON(&word); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": err.Error()})
			return
		}

		if ctx.Request.Header["Category"] == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "Category is required"})
			return
		}

		ctg := ctx.Request.Header["Category"][0]
		err := AddWord(word, cc, ctg)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		}
	}

	return gin.HandlerFunc(fn)
}

func CorrectWord(cc *mongo.Collection) gin.HandlerFunc {

	fn := func(ctx *gin.Context) {

		var word Word

		if err := ctx.ShouldBindJSON(&word); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": err.Error()})
			return
		}

		if ctx.Request.Header["Category"] == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "Category is required"})
			return
		}

		ctg := ctx.Request.Header["Category"][0]
		err := Corrector(word, cc, ctg)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		}
	}

	return gin.HandlerFunc(fn)
}

func EditWord(cc *mongo.Collection) gin.HandlerFunc {

	fn := func(ctx *gin.Context) {

		var body RequestDTO

		er := ctx.BindJSON(&body)

		if er != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": er.Error()})
			return
		}

		fmt.Println(body.New)

		if body.Word == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "word is required"})
			return
		}

		if ctx.Request.Header["Category"] == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "category is required"})
			return
		}

		if body.New == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "new word is required"})
			return
		}

		word := body.Word
		newValue := body.New
		ctg := ctx.Request.Header["Category"][0]
		err := Edit(word, newValue, cc, ctg)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		}
	}

	return gin.HandlerFunc(fn)
}

func DeleteWord(cc *mongo.Collection) gin.HandlerFunc {

	fn := func(ctx *gin.Context) {

		var body RequestDTO

		er := ctx.BindJSON(&body)
		if er != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": er.Error()})
			return
		}
		if body.Word == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "word is required"})
			return
		}
		if ctx.Request.Header["Category"] == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "Category is required"})
			return
		}

		ctg := ctx.Request.Header["Category"][0]
		err := Delete(body.Word, cc, ctg)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		}
	}

	return gin.HandlerFunc(fn)
}

type RequestDTO struct {
	Word string `json:"word"`
	New  string `json:"new"`
}
