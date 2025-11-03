package folders

type Folder struct {
	ID   uint    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;type:text"`
	Parent_ID uint `gorm:"type:int;foreignkey:ID"`
}