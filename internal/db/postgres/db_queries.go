package postgres

import (
	"KatodForAnod/siteGallery/internal/models"
	"errors"
	"fmt"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"html/template"
)

const addImage = `
	INSERT INTO img
	(id, user_id, img, tags)
	VALUES (DEFAULT,$1,$2,$3)
`

func (p postgreSQl) AddImage(data models.ImgMetaData) error {
	_, err := p.conn.Exec(addImage, data.UserId, string(data.Data), (*pq.StringArray)(&data.Tags))
	if err != nil {
		return fmt.Errorf("AddImage err:%s", err)
	}

	return nil
}

func (p postgreSQl) RemoveImage(id int64) error {
	panic("implement me")
}

const getImage = `
	SELECT id, img.img, tags
	FROM img
	WHERE id = $1
`

func (p postgreSQl) GetImage(id int64) (models.ImgMetaData, error) {
	rows, err := p.conn.Query(getImage, id)
	if err != nil {
		return models.ImgMetaData{}, fmt.Errorf("GetImage err: %s", err)
	}
	defer rows.Close()

	img := models.ImgMetaData{}
	if rows.Next() {
		var fileBody string
		err := rows.Scan(&img.Id, &fileBody, (*pq.StringArray)(&img.Tags))
		if err != nil {
			log.Error(err)
			return models.ImgMetaData{}, fmt.Errorf("GetImage err: %v", err)
		}
		img.Data = template.URL(fileBody)
		return img, nil
	} else {
		return models.ImgMetaData{}, fmt.Errorf("image not found")
	}
}

const getImages = `
	SELECT id, img.img, tags
	FROM img
	OFFSET $1
	LIMIT $2
`

func (p postgreSQl) GetImages(offSet, limit int64) ([]models.ImgMetaData, error) {
	rows, err := p.conn.Query(getImages, offSet, limit)
	if err != nil {
		return []models.ImgMetaData{}, fmt.Errorf("GetImages err: %s", err)
	}
	defer rows.Close()

	imges := make([]models.ImgMetaData, 0, 10)
	data := models.ImgMetaData{}
	for rows.Next() {
		var fileBody string
		err := rows.Scan(&data.Id, &fileBody, (*pq.StringArray)(&data.Tags))
		if err != nil {
			log.Error(err)
			continue
		}
		data.Data = template.URL(fileBody)
		imges = append(imges, data)
	}

	return imges, nil
}

const addUser = `
	INSERT INTO users
	(email, id, name, password)
	VALUES ($1, DEFAULT, $2, $3)
`

func (p postgreSQl) AddUser(user models.User) error {
	_, err := p.conn.Exec(addUser, user.Email, user.User, user.PassHash)
	if err != nil {
		return fmt.Errorf("AddUser err: %s", err)
	}

	return nil
}

const addUserRetId = `
	INSERT INTO users
	(email, id, name, password)
	VALUES ($1, DEFAULT, $2, $3)
	RETURNING id
`

func (p postgreSQl) AddUserRetId(user models.User) (int64, error) {
	rows, err := p.conn.Query(addUserRetId, user.Email, user.User, user.PassHash)
	if err != nil {
		return 0, fmt.Errorf("AddUserRetId err: %s", err)
	}

	if !rows.Next() {
		return 0, fmt.Errorf("AddUserRetId err")
	}
	var id int64
	if err := rows.Scan(&id); err != nil {
		return 0, fmt.Errorf("AddUserRetId err: %s", err)
	}

	return id, nil
}

const getUser = `
	SELECT email, id, name, password
	FROM users
	WHERE email = $1
`

func (p postgreSQl) GetUser(email string) (models.User, error) {
	rows, err := p.conn.Query(getUser, email)
	if err != nil {
		return models.User{}, fmt.Errorf("GetUser err:%s", err)
	}
	defer rows.Close()

	data := models.User{}
	for rows.Next() {
		err := rows.Scan(&data.Email, &data.Id, &data.User, &data.PassHash)
		if err != nil {
			log.Error(err)
			continue
		}

		return data, nil
	}

	return data, errors.New("user not found")
}
