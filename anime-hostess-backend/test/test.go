package main

import (
	"fmt"
	"github.com/juzeon/anime-hostess/include"
)

func main() {
	TestGetVideosOfDirectory()
}

func TestGetVideosOfDirectory() {
	videos, err := include.GetVideosOfDirectory("aa", "D:/Downloads/2017 结城友奈是勇者 SEASON 2")
	if err != nil {
		panic(err)
	}
	fmt.Println(videos)
}
