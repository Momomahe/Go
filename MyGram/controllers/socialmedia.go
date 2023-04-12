package controllers

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse2 struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
type CreateSocial struct {
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
}

// CreateSocialMedia godoc
// @Summary CreateSocialMedia a new socialmedia
// @Description CreateSocialMedia a new socialmedia with the given information
// @Tags socialmedia
// @Accept json
// @Produce json
// @Param socialmedia body CreateSocial true "The social media to create"
// @Success 201 {object} models.SocialMedia
// @Failure 400 {object} ErrorResponse2
// @Failure 500 {object} ErrorResponse2
// @Router /socialmedia/ [post]
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
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
			"message": "Social Media not found",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, socialmedia)
}

func GetAllSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmediaList := []models.SocialMedia{}

	err := db.WithContext(ctx).Find(&socialmediaList).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get Social Media data",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, socialmediaList)
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
