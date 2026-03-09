package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/JoaoCunha50/S3-Drive/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var input CreateUserRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser := User{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Role:     Member,
		Password: string(hash),
	}

	err = h.repo.CreateUser(context.Background(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *Handler) LoginUser(c *gin.Context) {
	var request LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.repo.LoginUser(context.Background(), request.Email, request.Username, request.Password)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.SignToken(user.ID.Hex(), user.Email, string(user.Role), user.Name, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "SERVER_SUCCESS", "success": true, "user": user, "token": token})
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.repo.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userResponse := UserResponse{
		ID:         user.ID.Hex(),
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email,
		Role:       user.Role,
		ProfilePic: user.ProfilePic,
	}

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}

func (h *Handler) GetUserByIdInt(c *gin.Context) {
	id := c.Param("id")
	_, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.repo.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userResponse := UserResponse{
		ID:         user.ID.Hex(),
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email,
		Role:       user.Role,
		ProfilePic: user.ProfilePic,
	}

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}
