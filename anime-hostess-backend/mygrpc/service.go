package mygrpc

import (
	"compress/flate"
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/juzeon/anime-hostess/include"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"io"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BulletService struct {
}

func (b BulletService) GetBullets(ctx context.Context, request *BulletRequest) (*BulletResponse, error) {
	resp, err := include.SharedHTTPClient.Do(include.CreateHTTPRequest("GET",
		"https://comment.bilibili.com/"+strconv.Itoa(int(request.Cid))+".xml", nil))
	err = include.CheckHTTPError(resp, err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(flate.NewReader(resp.Body))
	if err != nil {
		return nil, err
	}
	var bullets []*BulletEntity
	doc.Find("i d").Each(func(i int, selection *goquery.Selection) {
		p, _ := selection.Attr("p")
		arr := strings.Split(p, ",")
		if len(arr) != 9 {
			log.Println(p)
			return
		}
		t, _ := strconv.ParseFloat(arr[0], 32)
		typ, _ := strconv.Atoi(arr[1])
		size, _ := strconv.Atoi(arr[2])
		color, _ := strconv.Atoi(arr[3])
		sentAt, _ := strconv.Atoi(arr[4])
		poolType, _ := strconv.Atoi(arr[5])
		sender := arr[6]
		dmid, _ := strconv.Atoi(arr[7])
		level, _ := strconv.Atoi(arr[8])
		bullets = append(bullets, &BulletEntity{
			Time:     float32(t),
			Type:     int64(typ),
			Size:     int64(size),
			Color:    int64(color),
			SentAt:   int64(sentAt),
			PoolType: int64(poolType),
			Sender:   sender,
			Dmid:     int64(dmid),
			Level:    int64(level),
		})
	})
	return &BulletResponse{
		Data: bullets,
	}, nil
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
