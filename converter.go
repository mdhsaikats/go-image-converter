package imgconv

import (
	"fmt"
	"image"
	"image/color"
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
		return fmt.Errorf("failed to decode PNG: %w", err)
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

func JpegToPng(inputPath, outputPath string) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open the input file: %w", err)
	}
	defer inFile.Close()

	img, err := jpeg.Decode(inFile)
	if err != nil {
		return fmt.Errorf("failed to decode JPEG: %w", err)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create the output file: %w", err)
	}
	defer outFile.Close()

	if err := png.Encode(outFile, img); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	return nil
}

func RemoveBackground(inputPath, outputPath string, tolerance uint8) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	src, _, err := image.Decode(inFile)
	if err != nil {
		return err
	}

	bounds := src.Bounds()
	
	bg := color.RGBAModel.Convert(
		src.At(bounds.Min.X, bounds.Min.Y),
	).(color.RGBA)

	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			c := color.RGBAModel.Convert(
				src.At(x, y),
			).(color.RGBA)

			if similarColor(c, bg, tolerance) {
				// Make background transparent
				dst.SetRGBA(x, y, color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 0,
				})
			} else {
				dst.SetRGBA(x, y, c)
			}
		}
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, dst)
}

func similarColor(a, b color.RGBA, tolerance uint8) bool {
	diff := func(x, y uint8) uint8 {
		if x > y {
			return x - y
		}
		return y - x
	}

	return diff(a.R, b.R) <= tolerance &&
		diff(a.G, b.G) <= tolerance &&
		diff(a.B, b.B) <= tolerance
}