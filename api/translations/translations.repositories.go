package translations

import (
	"errors"

	"gorm.io/gorm"
)

type TranslationRepository struct {
	db *gorm.DB
}

func NewTranslationRepository(db *gorm.DB) *TranslationRepository {
	return &TranslationRepository{db: db}
}

func (r* TranslationRepository) CreateTranslation(translation *Translation) error {
	t := r.db.Create(translation)
	if t == nil {
		return errors.New("error creating translation")
	}

	return nil
}

func (r* TranslationRepository) GetTranslations() []Translation {
	var translations []Translation
	err := r.db.Find(&translations)

	if err == nil {
		return nil
	}

	return translations
}

func (r* TranslationRepository) GetTranslation(tag string, lang string) Translation {
	var translation Translation
	err := r.db.First(&translation, "tag = ? AND lang = ?", tag, lang)

	if err != nil {
		return Translation{}
	}

	return translation
}