package handlers

import (
	"github.com/tombull/teamdream/app/pkg/backup"
	"github.com/tombull/teamdream/app/pkg/web"
)

// ExportBackupZip returns a Zip file with all content
func ExportBackupZip() web.HandlerFunc {
	return func(c *web.Context) error {

		file, err := backup.Create(c)
		if err != nil {
			return c.Failure(err)
		}

		return c.Attachment("backup.zip", "application/zip", file.Bytes())
	}
}
