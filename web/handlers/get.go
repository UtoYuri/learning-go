package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Params struct {
	FieldA string `form:"field_a" binding:"required"`
	FieldB string `form:"field_b" binding:"required"`
}

type UriParams struct {
	FieldA string `uri:"field_a" binding:"required"`
	FieldB string `uri:"field_b" binding:"required"`
}

func GetByParams(c *gin.Context) {
	var params Params

	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": params,
	})
}

func GetByUriParams(c *gin.Context) {
	var params UriParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": params,
	})
}