package translations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TranslationHandler struct {
	repo *TranslationRepository
}

func NewTranslationHandler(repo *TranslationRepository) *TranslationHandler {
	return &TranslationHandler{repo: repo}
}

func (r *TranslationHandler) CreateTranslation(c* gin.Context) {
	var input CreateTranslation
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := Translation{
		Tag: input.Tag,
		Translation: input.Translation,
		Lang: input.Lang,
	}

	err = r.repo.CreateTranslation(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Translation created successfully"})
}

func (h *TranslationHandler) GetTranslations(c *gin.Context) {
	translations := h.repo.GetTranslations()
	if translations == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch translations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": translations})
}

func (h *TranslationHandler) GetTranslation(c *gin.Context) {
	tag := c.Query("tag")
	lang := c.Query("lang")

	if tag == "" || lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tag and lang are required"})
		return
	}

	tr := h.repo.GetTranslation(tag, lang)

	if (tr == Translation{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "translation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tr})
}
