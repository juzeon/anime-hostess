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
