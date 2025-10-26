package translations

type CreateTranslation struct {
	Tag   string `json:"tag"`
	Lang  string `json:"lang"`
	Translation string `json:"translation"`
}