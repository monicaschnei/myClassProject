package request

type AddAvailabilityRequest struct {
	Date        string `json:"date" validate:"required"`
	UserId      int64  `json:"user_id" validate:"required"`
	Username    string `json:"username" validate:"required,alphanum"`
	Start       string `json:"start" validate:"required"`
	EndTime     string `json:"end" validate:"required"`
	IsAvailable bool   `json:"is_available" validate:"required"`
}

type ListAvailabilityRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
