package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateUser(c *gin.Context)
	GetUsers(c *gin.Context)
}
