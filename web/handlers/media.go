package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-module/internal/models"
	"go-module/web/utils"
	"net/http"
)

type GetParams struct {
	ID uint `uri:"id" binding:"required"`
}

type GetBulkParams struct {
}

type CreateParams struct {
	Title string `form:"title" binding:"required"`
	Kind int `form:"kind"`
	Suffix string `form:"suffix"`
	Size uint64 `form:"size"`
	Url string `form:"url" binding:"required"`
}

type UpdateParams struct {
	Title string `form:"title" binding:"required"`
	Url string `form:"url" binding:"required"`
}

type PatchParams struct {
	Title string `form:"title"`
	Url string `form:"url"`
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

// @Summary Get bulk media
// @Tags Media
// @version 1.0
// @Accept application/json
// @Success 200 {array} models.MediaReadonly "Media infos"
// @Failure 400 {object} FailedResponse "Wrong params"
// @Failure 500 {object} FailedResponse "Other"
// @Router /media [get]
func GetMediaBulk(c *gin.Context) {
	var params GetBulkParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusBadRequest,
		})
		return
	}

	medias, err := models.GetMediaBulk()
	if err != nil {
		c.JSON(http.StatusInternalServerError, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.PlainBulk(medias))
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

// @Summary Patch media
// @Tags Media
// @version 1.0
// @Accept application/json
// @Param id path uint true "Media ID"
// @Param media body PatchParams true "Media info"
// @Success 200 {object} models.MediaReadonly "Media info"
// @Failure 400 {object} FailedResponse "Wrong params"
// @Failure 500 {object} FailedResponse "Other"
// @Router /media/{id} [patch]
func PatchMedia(c *gin.Context) {
	var getParams GetParams
	var patchParams PatchParams


	if err := c.ShouldBindUri(&getParams); err != nil {
		c.JSON(http.StatusBadRequest, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusBadRequest,
		})
		return
	}

	if err := c.ShouldBind(&patchParams); err != nil {
		c.JSON(http.StatusBadRequest, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusBadRequest,
		})
		return
	}


	media, err := models.GetMedia(getParams.ID)
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

	media.Title = utils.ValidOrDefault(patchParams.Title, media.Title).(string)
	media.Url = utils.ValidOrDefault(patchParams.Url, media.Url).(string)

	if err := media.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, media.Plain())
}

// @Summary Update media
// @Tags Media
// @version 1.0
// @Accept application/json
// @Param id path uint true "Media ID"
// @Param media body UpdateParams true "Media info"
// @Success 200 {object} models.MediaReadonly "Media info"
// @Failure 400 {object} FailedResponse "Wrong params"
// @Failure 500 {object} FailedResponse "Other"
// @Router /media/{id} [put]
func UpdateMedia(c *gin.Context) {
	var getParams GetParams
	var updateParams UpdateParams

	if err := c.ShouldBindUri(&getParams); err != nil {
		c.JSON(http.StatusBadRequest, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusBadRequest,
		})
		return
	}

	if err := c.ShouldBind(&updateParams); err != nil {
		c.JSON(http.StatusBadRequest, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusBadRequest,
		})
		return
	}


	media, err := models.GetMedia(getParams.ID)
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

	media.Title = updateParams.Title
	media.Url = updateParams.Url

	if err := media.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, FailedResponse{
			ErrMsg: err.Error(),
			ErrCode: http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, media.Plain())
}