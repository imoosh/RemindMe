package example

import (
	"RemindMe/model"
)

// file struct, 文件结构体
type ExaFile struct {
	models.Model
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	models.Model
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
