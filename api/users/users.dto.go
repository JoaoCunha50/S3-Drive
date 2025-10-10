package users

type UserResponseDTO struct {
	ID    int 
	Name  string
	Username string
	Email string
	ProfilePic string
	ProfileDescription string
}

type LoginRequest struct {
	Email *string
	Username *string
	Password string
}