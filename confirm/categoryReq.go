package confirm

type CreateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}