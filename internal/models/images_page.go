package models

import (
	"html/template"
)

type ImagesPageBody struct {
	ImagesArr       []ImgMetaData
	PagePrev        template.URL
	PageNext        template.URL
	ImagePlus       template.URL
	ImageBackground template.URL
}
