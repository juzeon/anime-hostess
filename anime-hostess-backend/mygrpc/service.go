package mygrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juzeon/anime-hostess/include"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BulletService struct {
}

func (b BulletService) GetAnime(ctx context.Context, request *AnimeRequest) (*AnimeResponse, error) {
	resp, err := include.SharedHTTPClient.Do(include.CreateHTTPRequest("GET",
		"https://api.bilibili.com/pgc/view/web/season?season_id="+strconv.Itoa(int(request.SeasonID)), nil))
	err = include.CheckHTTPError(resp, err)
	if err != nil {
		return nil, err
	}
	responseBytes, _ := io.ReadAll(resp.Body)
	var biliAnimeDetail include.BiliAnimeDetail
	_ = json.Unmarshal(responseBytes, &biliAnimeDetail)
	if biliAnimeDetail.Code != 0 {
		return nil, errors.New(biliAnimeDetail.Message)
	}
	var episodes []*EpisodeEntity
	for _, anime := range biliAnimeDetail.Result.Episodes {
		episodes = append(episodes, &EpisodeEntity{
			Cid:   int64(anime.Cid),
			Title: anime.Title + " " + anime.LongTitle,
		})
	}
	return &AnimeResponse{
		Data: episodes,
	}, nil
}

func (b BulletService) SearchAnime(ctx context.Context, request *SearchAnimeRequest) (*SearchAnimeResponse, error) {
	qs := url.Values{
		"search_type": {"media_bangumi"},
		"keyword":     {request.GetText()},
	}
	resp, err := include.SharedHTTPClient.Do(include.CreateHTTPRequest("GET",
		"https://api.bilibili.com/x/web-interface/search/type?"+qs.Encode(), nil))
	err = include.CheckHTTPError(resp, err)
	if err != nil {
		return nil, err
	}
	responseBytes, _ := io.ReadAll(resp.Body)
	biliSearchAnime := include.BiliSearchAnime{}
	_ = json.Unmarshal(responseBytes, &biliSearchAnime)
	var animeArr []*AnimeEntity
	if biliSearchAnime.Code != 0 {
		return nil, errors.New(biliSearchAnime.Message)
	}
	for _, anime := range biliSearchAnime.Data.Result {
		title := strings.ReplaceAll(anime.Title, "<em class=\"keyword\">", "")
		title = strings.ReplaceAll(title, "</em>", "")
		animeArr = append(animeArr, &AnimeEntity{
			SeasonID: int64(anime.SeasonId),
			Title:    title,
		})
	}
	return &SearchAnimeResponse{
		Data: animeArr,
	}, nil
}

var BulletClientImpl BulletClient

func ConnectServer() {
	ctx, cancel := include.GetGRPCContext()
	defer cancel()
	conn, err := grpc.DialContext(ctx,
		include.Config.GRPCServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             900 * time.Millisecond,
			PermitWithoutStream: true,
		}),
		grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	BulletClientImpl = NewBulletClient(conn)
	fmt.Println("Connected to grpc server " + include.Config.GRPCServer)
}
