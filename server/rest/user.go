package rest

import (
	"net/http"

	"github.com/Tonmoy404/project/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) CreateUser(ctx *gin.Context) {
	var user service.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	salt, err := bcrypt.GenerateFromPassword([]byte(user.Password), s.salt.SecretKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate salt"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword(salt, s.salt.SecretKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
	}

	user.Password = string(hashedPass)

	userRes := createUserRes{
		Username: user.Username,
		Email:    user.Email,
	}

	s.svc.CreateUser(&user)

	ctx.JSON(http.StatusOK, userRes)
}

func (s *Server) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := s.svc.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

	userRes := getUserRes{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	ctx.JSON(http.StatusOK, userRes)
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	existingUser, err := s.svc.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not find the user"})
		return
	}

	if existingUser == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user doesn't exist"})
		return
	}

	var updatedUser service.User
	err = ctx.ShouldBindJSON(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read the request body"})
		return
	}

	existingUser.Username = updatedUser.Username
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password

	err = s.svc.UpdateUser(userId, existingUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the user"})
		return
	}

	ctx.JSON(http.StatusOK, existingUser)

}

func (s *Server) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := s.svc.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
