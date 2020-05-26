package query

import "github.com/tombull/teamdream/app/models/dto"

type ListBlobs struct {
	Result []string
}

type GetBlobByKey struct {
	Key string

	Result *dto.Blob
}
