package stplib

import (
	"time"
)

type ImgInfo struct {
	FileName        string    // primary partion key
	CreatedAt       time.Time // time of a file created.
	OrgPath         string    // original file path
	ResizedFilePath string    // resized file path
	FileType        string    // file type of image: jpeg, gif, or png.
	SizeX           int       // image size x
	SizeY           int       // image size y
}
