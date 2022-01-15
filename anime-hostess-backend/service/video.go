package service

import (
	"github.com/juzeon/anime-hostess/include"
)

func VideoList() include.Result {
	series, err := include.GetAllSeries(false)
	if err != nil {
		return include.NewErrorResult(err)
	}
	return include.NewSuccessResult(series)
}
