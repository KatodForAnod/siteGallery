package view

import (
	"fmt"
	"net/http"
	"siteGallery/config"
)

func StartHttpServer(loadedConf config.Config) error {
	http.HandleFunc("/", Handlers{}.Default)

	fmt.Println("Server is listening...")
	addr := loadedConf.SvConfig.Host + ":" + loadedConf.SvConfig.Port
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}

	return nil
}
