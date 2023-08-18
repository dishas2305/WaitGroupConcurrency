package main

import (
	"channels/utils"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	starttime := time.Now()
	flag.Parse()
	imageURLs := flag.Args()

	// var wg sync.WaitGroup
	// for _, url := range imageURLs {
	// 	wg.Add(1)
	// 	go func(imageURL string) {
	// 		defer wg.Done()
	// 		imageData, err := utils.DownloadImage(imageURL)
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		imageConfig, err := utils.GetImageConfig(imageData)
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		scaledImage := utils.ResizeImage(imageData, imageConfig, 800, 0)
	// 		outputPath := "test_" + utils.GetFilename(imageURL)
	// 		if err := utils.SaveImageFile(scaledImage, outputPath); err != nil {
	// 			log.Println(err)
	// 		}
	// 	}(url)
	// }
	// wg.Wait()

	for _, url := range imageURLs {
		imageData, err := utils.DownloadImage(url)
		if err != nil {
			log.Println(err)
		}
		imageConfig, err := utils.GetImageConfig(imageData)
		if err != nil {
			log.Println(err)
		}
		scaledImage := utils.ResizeImage(imageData, imageConfig, 800, 0)
		outputPath := "test_" + utils.GetFilename(url)
		if err := utils.SaveImageFile(scaledImage, outputPath); err != nil {
			log.Println(err)
		}
	}
	endtime := time.Now()
	exec_time := endtime.Sub(starttime)
	fmt.Println("Execution time: ", exec_time)
}
