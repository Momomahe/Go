package controllers

import (
	"errors"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Post details for a given Id
// @Description post deatils of user corresponding to the input Id
// @Tags user
// @Accept json
// @produse json@Param models.User true "Register"
// @Success 201 (object) models.User
// @Router /user [post]
func Register(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// Login godoc
// @Summary Post details for a given Id
// @Description post deatils of user corresponding to the input Id
// @Tags user
// @Accept json
// @produse json@Param models.User true "Login"
// @Success 201 (object) models.User
// @Router /user [post]
func Login(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	password := user.Password

	err = db.Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
			"error":   "The email and password you provided are not associated with an existing account",
		})
		return
	}

	if !helpers.PasswordValid(user.Password, password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Email or Password",
			"error":   "You must fill in the correct email and password and have registered",
		})
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid password"))
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"your token": token,
	})
}
