package info

import (
	"github.com/JoaoCunha50/S3-Drive/backend/internal/translation"
	"github.com/JoaoCunha50/S3-Drive/backend/internal/user"
	"github.com/gin-gonic/gin"
)

type Service struct {
	UsersRepo        *user.Repository
	TranslationsRepo *translation.Repository
}

func NewService(usersRepo *user.Repository, translationsRepo *translation.Repository) *Service {
	return &Service{UsersRepo: usersRepo, TranslationsRepo: translationsRepo}
}

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetInfo(c *gin.Context) {
	currentUser, exists := c.Get("currentUser")
	if exists {
		currentUserMap := currentUser.(map[string]interface{})
		id := currentUserMap["id"].(string)
		user, err := h.Service.UsersRepo.GetUserByID(c.Request.Context(), id)
		if err == nil {
			c.JSON(200, gin.H{
				"user":               user,
				"translations":       h.Service.TranslationsRepo.GetTranslations,
				"defaultLanguage":    "en",
				"languagesSupported": []string{"en", "pt"},
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"translations":       []translation.Translation{},
		"defaultLanguage":    "en",
		"languagesSupported": []string{"en", "pt"},
	})
}

func RegisterRoutes(r *gin.RouterGroup, service *Service) {
	handler := NewHandler(service)
	r.GET("", handler.GetInfo)
}
