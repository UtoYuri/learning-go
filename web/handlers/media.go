package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-module/internal/models"
	"net/http"
)

type GetParams struct {
	ID uint `uri:"id" binding:"required"`
}

type CreateParams struct {
	Title string `form:"title" binding:"required"`
	Kind string `form:"kind"`
	Suffix string `form:"suffix"`
	Size uint64 `form:"size"`
	Url string `form:"url" binding:"required"`
}

// @Summary Get media by id
// @Tags Media
// @version 1.0
// @Accept application/json
// @Param id path uint true "Media ID"
// @Success 200 {object} models.MediaReadonly "Media info"
// @Failure 400 {object} FailedResponse "Wrong params"
// @Failure 404 {object} FailedResponse "Record not found"
// @Failure 500 {object} FailedResponse "Other"
// @Router /media/{id} [get]
func GetMedia(c *gin.Context) {
	var params GetParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusBadRequest,
		})
		return
	}

	media, err := models.GetMedia(params.ID)
	if err != nil {
		var errCode = http.StatusInternalServerError
		if gorm.IsRecordNotFoundError(err) {
			errCode = http.StatusNotFound
		}

		c.JSON(errCode, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: errCode,
		})
		return
	}

	c.JSON(http.StatusOK, media.Plain())
}



// @Summary Create media
// @Tags Media
// @version 1.0
// @Accept application/json
// @Param media body CreateParams true "Media info"
// @Success 200 {object} models.MediaReadonly "Media info"
// @Failure 400 {object} FailedResponse "Wrong params"
// @Failure 500 {object} FailedResponse "Other"
// @Router /media [post]
func CreateMedia(c *gin.Context) {
	var params CreateParams

	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusBadRequest,
		})
		return
	}

	media := &models.Media{
		Title: params.Title,
		Kind: params.Kind,
		Suffix: params.Suffix,
		Size: params.Size,
		Url: params.Url,
	}

	err := models.CreateMedia(media)
	if err != nil {
		c.JSON(http.StatusInternalServerError, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, media.Plain())
}