package service

import (
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/mygrpc"
)

func BulletSearch(text string) include.Result {
	ctx, cancel := include.GetGRPCContext()
	defer cancel()
	res, err := mygrpc.BulletClientImpl.SearchAnime(ctx, &mygrpc.SearchAnimeRequest{Text: text})
	if err != nil {
		return include.NewErrorResult("无法获取搜索结果：" + err.Error())
	}
	return include.NewSuccessResult(res.GetData())
}
func BulletAnime(seasonID int) include.Result {
	ctx, cancel := include.GetGRPCContext()
	defer cancel()
	res, err := mygrpc.BulletClientImpl.GetAnime(ctx, &mygrpc.AnimeRequest{SeasonID: int64(seasonID)})
	if err != nil {
		return include.NewErrorResult("无法获取剧集列表：" + err.Error())
	}
	return include.NewSuccessResult(res.GetData())
}
func BulletBullet(cid int) include.Result {
	ctx, cancel := include.GetGRPCContext()
	defer cancel()
	res, err := mygrpc.BulletClientImpl.GetBullets(ctx, &mygrpc.BulletRequest{Cid: int64(cid)})
	if err != nil {
		return include.NewErrorResult("无法获取弹幕：" + err.Error())
	}
	return include.NewSuccessResult(res.GetData())
}
