package model

type ImgMetaData struct {
	Id         int64
	FileName   string
	Tags       []string
	Data       string
	LoadByUser string
}

type TagToImages struct {
	Tag    string
	Images []ImgMetaData
}

type User struct {
	User     string
	Email    string
	PassHash string
}
