package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
		err := AddCorrectorFalse(word, cc, ctg)

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

		if ctx.Request.Header["Word"] == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "Word is required"})
			return
		}

		if ctx.Request.Header["Category"] == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "Category is required"})
			return
		}

		if ctx.Request.Header["NewValue"] == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err message": "NewValue is required"})
			return
		}

		word := ctx.Request.Header["Word"][0]
		newValue := ctx.Request.Header["NewValue"][0]
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
