package include

import (
	"fmt"
)

var ConfigPath *string

type Result struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func NewOKResult() Result {
	return Result{
		Status: true,
		Data:   nil,
	}
}
func NewSuccessResult(data interface{}) Result {
	return Result{
		Status: true,
		Data:   data,
	}
}
func NewErrorResult(data interface{}) Result {
	err, ok := data.(error)
	if ok {
		fmt.Println("#+v", err)
	}
	return Result{
		Status: false,
		Data:   data,
	}
}

type Video struct {
	SeriesName string `json:"seriesName"`
	Name       string `json:"name"`
	Path       string `json:"-"`
	Hash       string `json:"hash"`
}

type Series struct {
	Name   string  `json:"name"`
	Videos []Video `json:"videos"`
	Hash   string  `json:"hash"`
}
