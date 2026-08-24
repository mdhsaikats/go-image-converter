package imgconv

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

func PngToJpeg(inputPath, outputPath string, quality int) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w",err)
	}
	defer inFile.Close()
	img, err := png.Decode(inFile)
	if err != nil {
		fmt.Errorf("failed to decode PNG: %w",err)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create the output file: %w",err)
	}
	defer outFile.Close()

	option := &jpeg.Options{Quality: quality}
	err = jpeg.Encode(outFile,img,option)
	if err != nil {
		return fmt.Errorf("failed to encode JPEG: %w",err)
	}
	return nil
}
