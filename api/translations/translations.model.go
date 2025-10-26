package translations

type Translation struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
    Tag         string `gorm:"uniqueIndex:uniq_tag_lang;size:50;not null"`
    Lang        string `gorm:"uniqueIndex:uniq_tag_lang;size:5;not null"`
	Translation string `json:"translation" gorm:"type:text;not null"`
}