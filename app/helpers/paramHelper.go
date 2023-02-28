package helpers

// User
type SignUpRequestBody struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Wishlist
type CreateWishlistRequestBody struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Wishlist Item
type AddItemRequestBody struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Provider string `json:"provider"`
}

type UpdateItemRequestBody struct {
	Description string `json:"description"`
	Priority    string `json:"priority"`
}
