package model

type ImgMetaData struct {
	FileName   string
	Tags       []string
	Data       string
	LoadByUser string
}

type TagsToImages struct {
	Tag    string
	Images []ImgMetaData
}

type Users struct {
	User     string
	Email    string
	PassHash string
}
