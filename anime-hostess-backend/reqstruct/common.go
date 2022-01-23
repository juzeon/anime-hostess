package reqstruct

type BulletSearchRequest struct {
	Text string `uri:"text" binding:"required"`
}
type BulletAnimeRequest struct {
	SeasonID int `uri:"seasonID" binding:"required"`
}
type BulletBulletRequest struct {
	Cid int `uri:"cid"`
}

type UserSetProgressRequest struct {
	Hash string  `form:"hash" binding:"required"`
	Time float64 `form:"time" binding:"required,gte=0"`
}
type UserSetSearchTextRequest struct {
	Hash       string `form:"hash" binding:"required"`
	SearchText string `form:"searchText" binding:"required"`
}

type HashUriRequest struct {
	Hash string `uri:"hash" binding:"required"`
}
