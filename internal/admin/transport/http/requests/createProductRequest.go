package requests

type CreateProductRequest struct {
	CategoryID uint   `form:"category_id" validate:"required"`
	Title      string `form:"title" validate:"required"`
	//Slug string
	//Sku
	Status        string `form:"status"`
	Quantity      uint   `form:"quantity" validate:"required"`
	OriginalPrice uint   `form:"original_price" validate:"required"`
	SalePrice     uint   `form:"sale_price" validate:"required"`
	Description   string `form:"description" validate:"required"`
	Images        []string
}
