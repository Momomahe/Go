package controllers

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Create(&photo).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, photo)
}

func GetOnePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	err := db.WithContext(ctx).First(&photo, photoID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Photo not found",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Model(&photo).Where("id=?", photoID).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, Photo_Url: photo.Photo_Url}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	photo := models.Photo{}
	err := db.WithContext(ctx).First(&photo, photoID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Photo not found",
		})
		return
	}

	err = db.Delete(&models.Photo{}, photoID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Photo has been deleted",
	})
}
