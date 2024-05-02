package requests

import "mime/multipart"

type CreateCategoryRequest struct {
	Title  string                `form:"title" binding:"required"`
	Status string                `form:"status" `
	Image  *multipart.FileHeader `form:"image" binding:"required"`
}
