package include

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var allSeries []Series

func GetAllSeries(force bool) (series []Series, err error) {
	if len(allSeries) != 0 && force == false {
		return allSeries, nil
	}
	files, err := os.ReadDir(Config.VideoRoot)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if !fileInfo.IsDir() {
			continue
		}
		videos, err := GetVideosOfDirectory(fileInfo.Name(), path.Join(Config.VideoRoot, fileInfo.Name()))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(videos) == 0 {
			continue
		}
		singleSeries := Series{
			Name:   fileInfo.Name(),
			Videos: videos,
		}
		series = append(series, singleSeries)
	}
	allSeries = series
	return series, nil
}

var Hash2PathMap = map[string]string{}

func GetVideosOfDirectory(seriesName string, dir string) (videos []Video, err error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if fileInfo.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(fileInfo.Name()))
		if ext == "" {
			continue
		}
		if !StringInSlice(strings.Split(Config.VideoTypes, ","), ext[1:]) {
			continue
		}
		hash := MD5([]byte(path.Join(dir, fileInfo.Name())))
		videos = append(videos, Video{
			SeriesName: seriesName,
			Name:       fileInfo.Name(),
			Path:       path.Join(dir, fileInfo.Name()),
			Hash:       hash,
		})
		Hash2PathMap[hash] = path.Join(dir, fileInfo.Name())
	}
	return videos, nil
}

func StringInSlice(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
func MD5(data []byte) string {
	instance := md5.New()
	_, _ = io.Copy(instance, bytes.NewReader(data))
	return hex.EncodeToString(instance.Sum(nil))
}
func ValidateFields(validateFun func(valid *validation.Validation, ctx *gin.Context)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		valid := &validation.Validation{}
		validateFun(valid, ctx)
		if valid.HasErrors() {
			ctx.JSON(200, NewErrorResult(valid.Errors))
			ctx.Abort()
			return
		}
	}
}
func GetGRPCContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}
