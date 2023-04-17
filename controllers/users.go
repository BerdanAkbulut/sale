package controllers

import (
	"github.com/BerdanAkbulut/sale/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	err := models.GetUsers(&users)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUser(c *gin.Context) {
	var user models.User

	err := models.GetUser(&user, c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	err := models.CreateUser(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User

	err := models.GetUser(&user, c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.BindJSON(&user)

	err = models.UpdateUser(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user models.User

	err := models.GetUser(&user, c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	err = models.DeleteUser(&user, c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.Status(http.StatusNoContent)
	}
}
