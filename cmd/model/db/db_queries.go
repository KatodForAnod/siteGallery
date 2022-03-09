package db

import (
	"github.com/lib/pq"
	"html/template"
	"log"
	"siteGallery/cmd/model"
)

const addImage = `
	INSERT INTO img
	VALUES (DEFAULT,$1,$2)
`

func (p postgreSQl) AddImage(data model.ImgMetaData) error {
	_, err := p.conn.Exec(addImage, string(data.Data), (*pq.StringArray)(&data.Tags))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (p postgreSQl) RemoveImage(id int64) error {
	panic("implement me")
}

func (p postgreSQl) GetImage(id int64) (model.ImgMetaData, error) {
	panic("implement me")
}

const getImages = `
	SELECT id, img.img, tags
	FROM img
	OFFSET $1
	LIMIT $2
`

func (p postgreSQl) GetImages(offSet, limit int64) ([]model.ImgMetaData, error) {
	rows, err := p.conn.Query(getImages, offSet, limit)
	if err != nil {
		log.Println(err)
		return []model.ImgMetaData{}, err
	}
	defer rows.Close()

	imges := make([]model.ImgMetaData, 0, 10)
	data := model.ImgMetaData{}
	for rows.Next() {
		var fileBody string
		err := rows.Scan(&data.Id, &fileBody, (*pq.StringArray)(&data.Tags))
		if err != nil {
			log.Println(err)
			continue
		}
		data.Data = template.URL(fileBody)
		imges = append(imges, data)
	}

	return imges, nil
}

func (p postgreSQl) AddUser(user model.User) error {
	panic("implement me")
}

func (p postgreSQl) GetUser(email string) (model.User, error) {
	panic("implement me")
}