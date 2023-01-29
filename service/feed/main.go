package main

import (
	"log"
	douyin "nico_minidouyin/kitex_gen/douyin/feedservice"
)

func main() {
	svr := douyin.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
