package exeption

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()

	fmt.Println("Error LOG :", err != nil)

	if err == nil {
		return
	}

	switch true {
	case validationErrors(c, err):
		return
	case notFoundError(c, err):
		return
	case unAuthorizedError(c, err):
		return
	default:
		internalServerError(c, err)
	}
}

func validationErrors(c *gin.Context, err *gin.Error) bool {
	exception, ok := err.Err.(validator.ValidationErrors)
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   exception.Error(),
		})
		return true
	} else {
		return false
	}
}

func notFoundError(c *gin.Context, err *gin.Error) bool {
	exception, ok := err.Err.(*NotFoundError)
	if ok {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   http.StatusNotFound,
			"status": "NOT FOUND",
			"data":   exception.Err,
		})
		return true
	} else {
		return false
	}
}

func unAuthorizedError(c *gin.Context, err *gin.Error) bool {
	exception, ok := err.Err.(*UnAuthorizedError)
	if ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"status": err.Err.Error(),
			"data":   exception.Err,
		})
		return true
	} else {
		return false
	}
}

func internalServerError(c *gin.Context, err *gin.Error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":   http.StatusInternalServerError,
		"status": "INTERNAL SERVER ERROR",
		"data":   err,
	})
}
