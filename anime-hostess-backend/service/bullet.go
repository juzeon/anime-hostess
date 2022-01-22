package service

import (
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/mygrpc"
	"github.com/juzeon/anime-hostess/router/reqstruct"
)

func BulletSearch(request reqstruct.BulletSearchRequest) include.Result {
	ctx, cancel := include.GetGRPCContext()
	defer cancel()
	res, err := mygrpc.BulletClientImpl.SearchAnime(ctx, &mygrpc.SearchAnimeRequest{Text: request.Text})
	if err != nil {
		return include.NewErrorResult("无法获取搜索结果：" + err.Error())
	}
	return include.NewSuccessResult(res.GetData())
}
func BulletAnime(request reqstruct.BulletAnimeRequest) include.Result {
	ctx, cancel := include.GetGRPCContext()
	defer cancel()
	res, err := mygrpc.BulletClientImpl.GetAnime(ctx, &mygrpc.AnimeRequest{SeasonID: int64(request.SeasonID)})
	if err != nil {
		return include.NewErrorResult("无法获取剧集列表：" + err.Error())
	}
	return include.NewSuccessResult(res.GetData())
}
func BulletBullet(request reqstruct.BulletBulletRequest) include.Result {
	ctx, cancel := include.GetGRPCContext()
	defer cancel()
	res, err := mygrpc.BulletClientImpl.GetBullets(ctx, &mygrpc.BulletRequest{Cid: int64(request.Cid)})
	if err != nil {
		return include.NewErrorResult("无法获取弹幕：" + err.Error())
	}
	return include.NewSuccessResult(res.GetData())
}
