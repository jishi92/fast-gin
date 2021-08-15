package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	e "fast-gin/ecode"
	"fast-gin/library/util"
	"fast-gin/router/api"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]

func UploadImage(c *gin.Context) {
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		log.Println("err=", err)
		api.Response(c, http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	if image == nil {
		api.Response(c, http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	imageName := util.GetImageName(image.Filename)
	fullPath := util.GetImageFullPath()
	savePath := util.GetImagePath()
	src := fullPath + imageName

	if !util.CheckImageExt(imageName) || !util.CheckImageSize(file) {
		api.Response(c, http.StatusBadRequest, e.ErrorUploadCheckImageFormat, nil)
		return
	}

	err = util.CheckImage(fullPath)
	if err != nil {
		api.Response(c, http.StatusInternalServerError, e.ErrorUploadCheckImageFail, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		api.Response(c, http.StatusInternalServerError, e.ErrorUploadSaveImageFail, nil)
		return
	}

	api.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      util.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
