package utils

import (
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func DownloadImage(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	imgData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return imgData, nil
}

func GetImageConfig(imageData []byte) (image.Config, error) {
	config, _, err := image.DecodeConfig(strings.NewReader(string(imageData)))
	if err != nil {
		return image.Config{}, err
	}
	return config, err
}

func ResizeImage(imageData []byte, config image.Config, newWidth, newHeight uint) image.Image {
	img, _, _ := image.Decode(strings.NewReader(string(imageData)))
	scaledImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return scaledImg
}

func GetFilename(url string) string {
	splitUrl := strings.Split(url, "/")
	return splitUrl[len(splitUrl)-1]
}

func SaveImageFile(img image.Image, outputpath string) error {
	outputFile, err := os.Create(outputpath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		return err
	}

	return nil
}

func ProcessImage(url string) {
	imageData, err := DownloadImage(url)
	if err != nil {
		log.Println(err)
	}
	imageConfig, err := GetImageConfig(imageData)
	if err != nil {
		log.Println(err)
	}
	scaledImage := ResizeImage(imageData, imageConfig, 800, 0)
	outputPath := "test_" + GetFilename(url)
	if err := SaveImageFile(scaledImage, outputPath); err != nil {
		log.Println(err)
	}
}
