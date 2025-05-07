package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const PATH = "assets"

func UploadFile(file *multipart.FileHeader, path string) error {
	parts := strings.Split(path, "/")
	fileID := parts[1]
	dirPath := fmt.Sprintf("%s/%s", PATH, parts[0])

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0777); err != nil {
			return err
		}
	}

	filePath := fmt.Sprintf("%s/%s", dirPath, fileID)

	uploadedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadedFile.Close()

	// Using os.Create to open the file with appropriate permissions
	targetFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	// Copy file contents from uploadedFile to targetFile
	_, err = io.Copy(targetFile, uploadedFile)
	if err != nil {
		return err
	}

	return nil
}

func GetMimetype(f multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	return mimeType, nil
}

func GetExtensions(filename string) string {
	return strings.Split(filename, ".")[1]
}
