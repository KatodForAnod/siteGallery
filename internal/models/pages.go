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

type ViewImagePage struct {
	Image           ImgMetaData
	ImageBackground template.URL
}

type LoginRegistrationPage struct {
	ActionLink string
	ActionName string

	ImageBackground template.URL
}
