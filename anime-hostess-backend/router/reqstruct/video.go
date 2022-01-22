package reqstruct

type VideoByHashRequest struct {
	Hash string `uri:"hash" binding:"required"`
}
