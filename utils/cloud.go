package utils

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadCloudinary(file *multipart.FileHeader, folderName string) (string, error) {
	if file == nil {
		return "", nil
	}
	fileData, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	cldService, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	uploadParams := uploader.UploadParams{
		Folder:         folderName,
		UniqueFilename: true,
		Overwrite:      true,
	}

	resp, err := cldService.Upload.Upload(ctx, fileData, uploadParams)
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
