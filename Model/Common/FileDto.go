package Common

import "mime/multipart"

type FileDto struct {
	Data string                `form:"data"`
	File *multipart.FileHeader `form:"file"`
}
