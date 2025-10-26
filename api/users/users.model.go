package users

type User struct {
	ID         int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string  `gorm:"not null" json:"name"`
    Username   string  `gorm:"size:50;uniqueIndex;not null" json:"username"`
    Email      string  `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password   string  `gorm:"not null" json:"password"`
	ProfilePic *string `gorm:"type:text" json:"profilePic,omitempty"`
}