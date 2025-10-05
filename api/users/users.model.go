package users

type User struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	ProfilePic string `gorm:"null" json:"profilePic"`
}