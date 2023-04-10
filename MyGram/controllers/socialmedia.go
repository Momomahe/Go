package controllers

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmedia := models.SocialMedia{}

	err := ctx.ShouldBindJSON(&socialmedia)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Create(&socialmedia).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, socialmedia)
}

func GetOne(ctx *gin.Context) {
	db := database.GetDB()
	socialmedia := models.SocialMedia{}
	socialmediaID, _ := strconv.Atoi(ctx.Param("socialmediaID"))

	err := db.WithContext(ctx).First(&socialmedia, socialmediaID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Product not found",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, socialmedia)
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmedia := models.SocialMedia{}
	socialmediaID, _ := strconv.Atoi(ctx.Param("socialmediaID"))

	err := ctx.ShouldBindJSON(&socialmedia)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Model(&socialmedia).Where("id=?", socialmediaID).Updates(models.SocialMedia{Name: socialmedia.Name, Social_Media_Url: socialmedia.Social_Media_Url}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, socialmedia)
}

func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmediaID, _ := strconv.Atoi(ctx.Param("socialmediaID"))

	socialmedia := models.SocialMedia{}
	err := db.WithContext(ctx).First(&socialmedia, socialmediaID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Social media not found",
		})
		return
	}

	err = db.Delete(&models.SocialMedia{}, socialmediaID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Social media has been deleted",
	})
}
