package Model

import "mime/multipart"

type (
	S3ObjectRequest struct {
		File   *multipart.FileHeader
		Bucket string
		Key    string
	}

	S3UrlRequest struct {
		Bucket string
		Key    string
	}
)
