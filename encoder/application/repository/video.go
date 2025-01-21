package repository

import (
	"encoder/domain"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDB struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDB {
	return &VideoRepositoryDB{Db: db}
}

func (repository VideoRepositoryDB) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	if err := repository.Db.Create(video).Error; err != nil {
		return nil, err
	}

	return video, nil
}

func (repository VideoRepositoryDB) Find(id string) (*domain.Video, error) {
	var video domain.Video

	repository.Db.First(&video, "id = ?", id)

	if video.ID == "" {
		return nil, fmt.Errorf("Video não encontrado")
	}

	return &video, nil
}
