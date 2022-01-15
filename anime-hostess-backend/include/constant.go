package include

import "flag"

var ConfigPath string

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
	return Result{
		Status: false,
		Data:   data,
	}
}

func init() {
	flag.Parse()
	ConfigPath = *flag.String("c", "", "ConfigPath path")
}
