package files

type File struct {
	ID   uint      `gorm:"primaryKey;autoIncrement"`
	Name string    `gorm:"not null;type:text"`
	S3_Key string  `gorm:"not null;type:text"`  //Caminho para o binário no MinIO
	Path string    `gorm:"not null;type:text"`  //PATH real caso haja atualizações de pastas
	Size uint      `gorm:"not null;type:int"`
	Parent_ID uint `gorm:"not null;type:int;foreignkey:ID"`
}