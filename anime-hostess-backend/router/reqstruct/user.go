package reqstruct

type UserGetProgressRequest struct {
	Hash string `uri:"hash" binding:"required"`
}
type UserSetProgressRequest struct {
	Hash string  `form:"hash" binding:"required"`
	Time float64 `form:"time" binding:"required,gte=0"`
}
