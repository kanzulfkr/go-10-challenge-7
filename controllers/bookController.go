package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID  int    `json:"book_id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Desc    string `json:"desc"`
}

var BookLibrary = map[int]Book{}

var ID int

func CreateBook(ctx *gin.Context) {
	var newBook Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ID++
	newBook.BookID = ID
	BookLibrary[ID] = newBook

	ctx.JSON(http.StatusCreated, gin.H{
		"Book": newBook,
	})
}

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"book": BookLibrary,
	})
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	convertedbookID, err := strconv.Atoi(bookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Wrong Parameter Type, Must Be Integer",
			"error_message": fmt.Sprintf("%s is not integer", bookID),
		})
		return
	}

	_, exist := BookLibrary[convertedbookID]
	if exist {
		ctx.JSON(http.StatusOK, gin.H{
			"book": BookLibrary[convertedbookID],
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("%d not found", convertedbookID),
		})
		return
	}

}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	convertedbookID, err := strconv.Atoi(bookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Wrong Parameter Type, Must Be Integer",
			"error_message": fmt.Sprintf("%s is not integer", bookID),
		})
		return
	}

	_, exist := BookLibrary[convertedbookID]
	if exist {
		updatedBook.BookID = convertedbookID
		BookLibrary[convertedbookID] = updatedBook
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d successfully updated", convertedbookID),
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("%d not found", convertedbookID),
		})
		return
	}

}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	convertedbookID, err := strconv.Atoi(bookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Wrong Parameter Type, Must Be Integer",
			"error_message": fmt.Sprintf("%s is not integer", bookID),
		})
		return
	}

	_, exist := BookLibrary[convertedbookID]
	if exist {
		delete(BookLibrary, convertedbookID)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d successfully deleted", convertedbookID),
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("%d not found", convertedbookID),
		})
		return
	}

}
