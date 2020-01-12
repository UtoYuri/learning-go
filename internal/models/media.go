package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go-module/pkg/database"
)

const (
	StatusNormal = 0
	StatusDeleted = 1
)

type Media struct {
	gorm.Model
	Title string `gorm:"size:255;not null;"`
	Kind int
	Suffix string
	Size uint64
	Url string `gorm:"size:1000;not null;"`
	Status int `gorm:"default:'0'"`
}

type MediaReadonly struct {
	ID uint `form:"id"`
	Title string `form:"title"`
	Kind int `form:"kind"`
	Suffix string `form:"suffix"`
	Size uint64 `form:"size"`
	Url string `form:"url"`
	Status int `form:"status"`
	CreatedAt int64 `form:"created_at"`
	UpdatedAt int64 `form:"updated_at"`
}

func GetMedia(id uint) (*Media, error) {
	var media Media
	result := database.DB.First(&media, id)

	if media.Status != StatusNormal {
		return nil, errors.New("media forbidden")
	}

	return &media, result.Error
}

func GetMediaBulk() ([]*Media, error) {
	var medias []*Media

	result := database.DB.Where("status = ?", StatusNormal).Find(&medias)

	return medias, result.Error
}

func CreateMedia(media *Media) error {
	result := database.DB.Create(media)
	return result.Error
}

func (media *Media) Update() error {
	if media.Status != StatusNormal {
		return errors.New("media forbidden")
	}

	result := database.DB.Model(media).Updates(*media)
	return result.Error
}

func (media *Media) Delete() error {
	if media.Status != StatusNormal {
		return nil
	}

	media.Status = StatusDeleted
	result := database.DB.Model(media).Updates(*media)
	return result.Error
}

func (media *Media) Plain() MediaReadonly {
	return MediaReadonly{
		ID: media.ID,
		Title: media.Title,
		Kind: media.Kind,
		Suffix: media.Suffix,
		Size: media.Size,
		Url: media.Url,
		Status: media.Status,
		CreatedAt: media.CreatedAt.Unix(),
		UpdatedAt: media.UpdatedAt.Unix(),
	}
}

func PlainBulk(medias []*Media) []MediaReadonly {
	mediasReadonly := make([]MediaReadonly, len(medias))

	for _, media := range medias {
		mediasReadonly = append(mediasReadonly, MediaReadonly{
			ID:        media.ID,
			Title:     media.Title,
			Kind:      media.Kind,
			Suffix:    media.Suffix,
			Size:      media.Size,
			Url:       media.Url,
			Status:    media.Status,
			CreatedAt: media.CreatedAt.Unix(),
			UpdatedAt: media.UpdatedAt.Unix(),
		})
	}

	return mediasReadonly
}