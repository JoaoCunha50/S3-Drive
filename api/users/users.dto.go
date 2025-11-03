package users

type UserResponseDTO struct {
	Id    uint `json:"id"`
	Name  string `json:"name"` 
	Username string `json:"username"`
	Email string `json:"email"`
	ProfilePic *string `json:"profilePic,omitempty"`
}

type LoginRequest struct {
    Email    *string `json:"email"`
    Username *string `json:"username"`
    Password string `json:"password"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}