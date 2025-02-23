package response

type UserLoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
