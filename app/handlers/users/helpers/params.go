package helpers

type SignUpRequestBody struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
