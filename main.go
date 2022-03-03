package main

import (
	"siteGallery/config"
	"siteGallery/view"
)

func main() {
	conf := config.Config{}
	view.StartHttpServer(conf)
}
