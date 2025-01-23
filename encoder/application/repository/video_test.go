package repository_test

import (
	"encoder/application/repository"
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&domain.Video{})
	require.NoError(t, err)

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now().UTC()

	repo := repository.VideoRepositoryDB{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.Nil(t, err)
	require.NotEmpty(t, v.ID)
	require.Equal(t, v.ID, video.ID)
	require.Equal(t, v.FilePath, video.FilePath)
	require.Equal(t, v.CreatedAt, video.CreatedAt)

}
