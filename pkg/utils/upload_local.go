package utils

import (
	"go-fiber-api/pkg/dto"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type File struct {
	Name      string `json:"name"`
	BasePath  string `json:"base_path"`
	Dir       string `json:"dir"`
	Extension string `json:"extension"`
}

func Upload(files []*multipart.FileHeader) ([]*dto.FileStoreUploadResponse, error) {
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, err
	}

	var filesInfo []*dto.FileStoreUploadResponse

	for _, file := range files {
		changeFile, err := GenerateRandomFilename(file.Filename)
		if err != nil {
			return nil, err
		}
		dst := filepath.Join(uploadDir, changeFile)

		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		dst_file, err := os.Create(dst)
		if err != nil {
			return nil, err
		}
		defer dst_file.Close()

		buffer := make([]byte, file.Size)
		n, err := io.ReadFull(src, buffer)
		if err != nil {
			return nil, err
		}

		if err := os.WriteFile(dst, buffer[:n], 0644); err != nil {
			return nil, err
		}

		filesInfo = append(filesInfo, &dto.FileStoreUploadResponse{
			Name:      changeFile,
			BasePath:  uploadDir,
			Extension: filepath.Ext(changeFile),
		})
	}

	return filesInfo, nil
}
