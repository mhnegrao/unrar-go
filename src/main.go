package main

import (
	"archive/zip"

	"fmt"
	//"io"
	"log"
	"os"

	_"github.com/gen2brain/go-unarr"
	//"path/filepath"
)

// func main() {
// 	zipReader, _ := zip.OpenReader("arqcompact.rar")
// 	for _, file := range zipReader.Reader.File {

// 		zippedFile, err := file.Open()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer zippedFile.Close()

// 		targetDir := "./"
// 		extractedFilePath := filepath.Join(
// 			targetDir,
// 			file.Name,
// 		)

// 		if file.FileInfo().IsDir() {
// 			log.Println("Directory Created:", extractedFilePath)
// 			os.MkdirAll(extractedFilePath, file.Mode())
// 		} else {
// 			log.Println("File extracted:", file.Name)

// 			outputFile, err := os.OpenFile(
// 				extractedFilePath,
// 				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
// 				file.Mode(),
// 			)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			defer outputFile.Close()

// 			_, err = io.Copy(outputFile, zippedFile)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 	}
// }

func main1() {

	// a, err := unarr.NewArchive("arqcompact.zip")
	// if err != nil {
	// 	panic(err)
	// }
	// defer a.Close()
	// erro,_ := a.Extract("src")
	// if erro != nil {
	// 	panic(err)
	// }
	read, err := zip.OpenReader("go.rar")
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()
	for _, file := range read.File {
		if err := listFiles(file); err != nil {
			log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
		}
	}
}
func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()
	fmt.Fprintf(os.Stdout, "%s:", file.Name)
	if err != nil {
		msg := "failed to read zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	fmt.Println()
	return nil
}
