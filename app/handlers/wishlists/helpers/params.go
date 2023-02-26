package helpers

type CreateWishlistRequestBody struct {
	Name		string 	`json:"name"`
	Type		string 	`json:"type"`
	Description	string 	`json:"description"`
}

type AddItemRequestBody struct {
	Name		string 	`json:"name"`
	Url			string 	`json:"url"`
	Provider	string 	`json:"provider"`
}

type UpdateItemRequestBody struct {
	Description	string 	`json:"description"`
	Priority	string 	`json:"priority"`
}
