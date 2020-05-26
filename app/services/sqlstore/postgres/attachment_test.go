package postgres_test

import (
	"context"
	"testing"

	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/cmd"

	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/bus"
)

func TestUploadImage(t *testing.T) {
	ctx := SetupDatabaseTest(t)
	defer TeardownDatabaseTest()

	bus.AddHandler(func(ctx context.Context, c *cmd.StoreBlob) error {
		return nil
	})

	uploadImage := &cmd.UploadImage{
		Image: &models.ImageUpload{
			Upload: &models.ImageUploadData{
				Content:     []byte("Hello World"),
				ContentType: "text/plain",
			},
		},
		Folder: "avatars",
	}
	err := bus.Dispatch(ctx, uploadImage)
	Expect(err).IsNil()
	Expect(uploadImage.Image.BlobKey).ContainsSubstring("avatars/")
	Expect(uploadImage.Image.BlobKey).HasLen(73)
}

func TestUploadImage_NoContent(t *testing.T) {
	ctx := SetupDatabaseTest(t)
	defer TeardownDatabaseTest()

	bus.AddHandler(func(ctx context.Context, c *cmd.StoreBlob) error {
		return nil
	})

	uploadImage := &cmd.UploadImage{
		Image: &models.ImageUpload{
			Upload: &models.ImageUploadData{
				Content: []byte(""),
			},
		},
		Folder: "avatars",
	}
	err := bus.Dispatch(ctx, uploadImage)
	Expect(err).IsNil()
	Expect(uploadImage.Image.BlobKey).Equals("")
}

func TestUploadMultipleImages(t *testing.T) {
	ctx := SetupDatabaseTest(t)
	defer TeardownDatabaseTest()

	bus.AddHandler(func(ctx context.Context, c *cmd.StoreBlob) error {
		return nil
	})

	uploadImages := &cmd.UploadImages{
		Images: []*models.ImageUpload{
			&models.ImageUpload{
				Upload: &models.ImageUploadData{
					Content:     []byte("Hello World 1"),
					ContentType: "text/plain",
				},
			},
			&models.ImageUpload{
				Upload: &models.ImageUploadData{
					Content:     []byte("Hello World 2"),
					ContentType: "text/plain",
				},
			},
		},
		Folder: "avatars",
	}
	err := bus.Dispatch(ctx, uploadImages)
	Expect(err).IsNil()

	Expect(uploadImages.Images[0].BlobKey).ContainsSubstring("avatars/")
	Expect(uploadImages.Images[0].BlobKey).HasLen(73)

	Expect(uploadImages.Images[1].BlobKey).ContainsSubstring("avatars/")
	Expect(uploadImages.Images[1].BlobKey).HasLen(73)
}
