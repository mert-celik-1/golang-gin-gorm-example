package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createResponse(res interface{}, e error, funcName string, c *gin.Context) {

	r := c.Request
	w := c.Writer
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	status := http.StatusOK
	if e != nil {
		log.Panic(ctx, "", funcName, e) // TODO : Log kütüphanesi gelecek
		status = http.StatusBadRequest
	}

	c.JSON(status, res)
}
