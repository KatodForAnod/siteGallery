package models

import "html/template"

type ImgMetaData struct {
	Id         int64
	FileName   string
	Tags       []string
	Data       template.URL
	LoadByUser string
}

type TagToImages struct {
	Tag    string
	Images []ImgMetaData
}

type User struct {
	Id       int64
	User     string
	Email    string
	PassHash string
}
