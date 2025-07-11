package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"cadre-management/models"
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/pkg/logging"
	"cadre-management/pkg/upload"
	"cadre-management/pkg/utils"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CHECK_TOKEN_FAIL, nil)
		return
	}

	jwtClaims, ok := claims.(*utils.Claims)
	if !ok {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CHECK_TOKEN_FAIL, nil)
		return
	}
	userID := jwtClaims.UserID

	// 更新数据库记录
	err = models.UpdateCadreInfoModPhotoURL(userID, imageName)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
