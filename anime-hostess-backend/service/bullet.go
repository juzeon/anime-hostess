package service

import (
	"fmt"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/mygrpc"
)

func BulletSearch(text string) include.Result {
	ctx, cancel := include.GetGRPCContext()
	defer cancel()
	res, err := mygrpc.BulletClientImpl.SearchAnime(ctx, &mygrpc.SearchAnimeRequest{Text: text})
	if err != nil {
		return include.NewErrorResult("无法获取搜索结果：" + fmt.Sprintf("%+v", err))
	}
	return include.NewSuccessResult(res.GetData())
}
func BulletAnime(seasonID int) include.Result {
	return include.NewOKResult()
}
