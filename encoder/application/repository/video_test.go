package repository

import (
	"encoder/application/repository"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repository.VideoRepositoryDB{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.Nil(t, err)
	require.NotEmpty(t, v.ID)
	require.Equal(t, v.ID, video.ID)
	require.Equal(t, v.FilePath, video.FilePath)
	require.Equal(t, v.CreatedAt, video.CreatedAt)

}
