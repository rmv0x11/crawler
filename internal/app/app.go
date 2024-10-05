package app

import (
	"github.com/rmv0x11/crawler/internal/crawler"
	"github.com/rmv0x11/crawler/logger"
)

func Run() {
	logger.Init("debug")

	crawler.Start()

	select {}
}
