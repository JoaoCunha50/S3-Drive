package users

type UserResponseDTO struct {
	id    int 
	name  string
	username string
	email string
	profilePic string
}

type LoginRequest struct {
	email *string
	username *string
	password string
}