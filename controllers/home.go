package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, "Home route is working!")
}
