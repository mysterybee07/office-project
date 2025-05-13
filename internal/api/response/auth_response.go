package response

type LoginResponse struct {
	Email        string `json:"email" example:"user@example.com"`
	Role         string `json:"role" example:"admin"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message" example:"Login successful"`
}
