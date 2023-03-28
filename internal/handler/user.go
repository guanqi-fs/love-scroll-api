package handler

import (
	"love-scroll-api/internal/config"
	"love-scroll-api/internal/errorcode"
	"love-scroll-api/internal/response"
	"love-scroll-api/internal/service"
	"love-scroll-api/pkg/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := c.PostForm("role")

	db, _ := c.MustGet("db").(*database.DB)
	user, err := service.CreateUser(username, password, role, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(errorcode.DbErr, err, nil))
		return
	}

	secret := config.GetConfig().JWT.Secret
	token, err := service.GenerateToken(user, secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(errorcode.GenerateTokenErr, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{
		"token": token,
	}))
}

func LoginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	db, _ := c.MustGet("db").(*database.DB)
	user, err := service.CheckUserPassword(username, password, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewResponse(errorcode.Unauthorized, err, nil))
		return
	}

	secret := config.GetConfig().JWT.Secret
	token, err := service.GenerateToken(user, secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(errorcode.GenerateTokenErr, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{
		"token": token,
	}))
}

func GetUserHandler(c *gin.Context) {
	username := c.Param("username")

	db, _ := c.MustGet("db").(*database.DB)
	user, err := service.GetUser(username, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(errorcode.DbErr, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(user))
}

func UpdateUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(errorcode.BadRequestArgs, err, nil))
		return
	}

	role := c.PostForm("role")

	db, _ := c.MustGet("db").(*database.DB)
	user, err := service.GetUserByID(uint(userID), db)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(errorcode.DbErr, err, nil))
		return
	}

	user.Role = role
	err = service.UpdateUser(user, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(errorcode.DbErr, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{"message": "user updated"}))
}

func DeleteUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(errorcode.BadRequestArgs, err, nil))
		return
	}

	db, _ := c.MustGet("db").(*database.DB)
	err = service.DeleteUser(uint(userID), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(errorcode.DbErr, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{"message": "user deleted"}))
}

func ListUsersHandler(c *gin.Context) {
	db, _ := c.MustGet("db").(*database.DB)
	users, err := service.ListUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(errorcode.DbErr, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{
		"users": users,
	}))
}
