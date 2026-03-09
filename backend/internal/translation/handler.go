package translation

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateTranslation(c *gin.Context) {
	var input CreateTranslationRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := Translation{
		Tag:         input.Tag,
		Translation: input.Translation,
		Lang:        input.Lang,
	}

	err = h.repo.CreateTranslation(context.Background(), &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Translation created successfully"})
}

func (h *Handler) GetTranslations(c *gin.Context) {
	translations, err := h.repo.GetTranslations(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch translations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": translations})
}

func (h *Handler) GetTranslation(c *gin.Context) {
	tag := c.Query("tag")
	lang := c.Query("lang")

	if tag == "" || lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tag and lang are required"})
		return
	}

	tr, err := h.repo.GetTranslation(context.Background(), tag, lang)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "translation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tr})
}
