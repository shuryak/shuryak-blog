package dto

type GetManyRequest struct {
	Offset uint32 `form:"offset" json:"offset" binding:"min=0,required" example:"10"`
	Count  uint32 `form:"count" json:"count" binding:"min=0,required" example:"5"`
}
