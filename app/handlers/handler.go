package handlers

import (
	"net/http"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

func CreateUser() {

}

func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := services.GetByID(id)
	if err != nil {
		// gestion de l'erreur
		c.JSON(http.StatusInternalServerError, err)
	}

	if user == nil {
		c.JSON(http.StatusNotFound, nil)
	}
	c.JSON(http.StatusOK, user)

}
