package db

import (
	"github.com/lib/pq"
	"log"
	"siteGallery/model"
)

func (p postgreSQl) AddImage(data model.ImgMetaData) error {
	panic("implement me")
}

func (p postgreSQl) RemoveImage(id int64) error {
	panic("implement me")
}

func (p postgreSQl) GetImage(id int64) (model.ImgMetaData, error) {
	panic("implement me")
}

const getImages = `
	SELECT id, img_data, tags, load_by_user
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
		err := rows.Scan(&data.Id, &data.Data, (*pq.StringArray)(&data.Tags), &data.LoadByUser)
		if err != nil {
			log.Println(err)
			continue
		}
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
