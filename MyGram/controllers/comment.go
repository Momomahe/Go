package controllers

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	photo := models.Photo{}

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).First(&photo, comment.PhotoID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Photo not found",
		})
		return
	}

	err = db.WithContext(ctx).Create(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func GetOneComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	err := db.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Comment not found",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Model(&comment).Where("id=?", commentID).Updates(models.Comment{Message: comment.Message}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	comment := models.Comment{}
	err := db.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Comment not found",
		})
		return
	}

	err = db.Delete(&models.Comment{}, commentID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comment has been deleted",
	})
}
func GetAllComment(ctx *gin.Context) {
	db := database.GetDB()
	commentList := []models.Comment{}

	err := db.WithContext(ctx).Find(&commentList).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get Comment data",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, commentList)
}
