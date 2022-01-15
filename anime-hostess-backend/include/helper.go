package include

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

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
		videos = append(videos, Video{
			SeriesName: seriesName,
			Name:       fileInfo.Name(),
			Path:       path.Join(dir, fileInfo.Name()),
		})
	}
	return videos, nil
}

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
