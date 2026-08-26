package imgconv

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestPngToJpeg(t *testing.T) {
	tempDir := t.TempDir()

	inputPath := filepath.Join(tempDir, "test.png")
	outputPath := filepath.Join(tempDir, "test.jpg")

	createDummyPNG(t, inputPath)

	err := PngToJpeg(inputPath, outputPath, 90)
	if err != nil {
		t.Fatalf("PngToJpeg failed: %v", err)
	}

	info, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("Expected output file was not created: %v", err)
	}

	if info.Size() == 0 {
		t.Errorf("Output file is empty")
	}

	// Make sure the output is actually a valid JPEG.
	file, err := os.Open(outputPath)
	if err != nil {
		t.Fatalf("Failed to open output JPEG: %v", err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		t.Fatalf("Output is not a valid JPEG: %v", err)
	}

	if img.Bounds().Dx() != 1 || img.Bounds().Dy() != 1 {
		t.Errorf(
			"Expected image size 1x1, got %dx%d",
			img.Bounds().Dx(),
			img.Bounds().Dy(),
		)
	}
}

func TestJpegToPng(t *testing.T) {
	tempDir := t.TempDir()

	inputPath := filepath.Join(tempDir, "test.jpg")
	outputPath := filepath.Join(tempDir, "test.png")

	createDummyJPEG(t, inputPath)

	err := JpegToPng(inputPath, outputPath)
	if err != nil {
		t.Fatalf("JpegToPng failed: %v", err)
	}

	info, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("Expected output file was not created: %v", err)
	}

	if info.Size() == 0 {
		t.Errorf("Output file is empty")
	}

	// Make sure the output is actually a valid PNG.
	file, err := os.Open(outputPath)
	if err != nil {
		t.Fatalf("Failed to open output PNG: %v", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatalf("Output is not a valid PNG: %v", err)
	}

	if img.Bounds().Dx() != 1 || img.Bounds().Dy() != 1 {
		t.Errorf(
			"Expected image size 1x1, got %dx%d",
			img.Bounds().Dx(),
			img.Bounds().Dy(),
		)
	}
}

func TestRemoveBackground(t *testing.T) {
	tempDir := t.TempDir()

	inputPath := filepath.Join(tempDir, "bg_test.png")
	outputPath := filepath.Join(tempDir, "bg_removed.png")

	// Create a 2x2 image:
	//
	// (0,0): White
	// (1,0): Off-white
	// (0,1): Red
	// (1,1): White
	//
	// White is our background.
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	img.Set(0, 0, color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})

	img.Set(1, 0, color.RGBA{
		R: 250,
		G: 250,
		B: 250,
		A: 255,
	})

	img.Set(0, 1, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})

	img.Set(1, 1, color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})

	file, err := os.Create(inputPath)
	if err != nil {
		t.Fatalf("Failed to create input PNG: %v", err)
	}

	if err := png.Encode(file, img); err != nil {
		file.Close()
		t.Fatalf("Failed to encode input PNG: %v", err)
	}

	file.Close()

	// Remove background.
	err = RemoveBackground(inputPath, outputPath, 10)
	if err != nil {
		t.Fatalf("RemoveBackground failed: %v", err)
	}

	// Open output.
	outFile, err := os.Open(outputPath)
	if err != nil {
		t.Fatalf("Failed to open output image: %v", err)
	}
	defer outFile.Close()

	outImg, err := png.Decode(outFile)
	if err != nil {
		t.Fatalf("Failed to decode output image: %v", err)
	}

	// Check top-left background.
	c00 := color.RGBAModel.Convert(
		outImg.At(0, 0),
	).(color.RGBA)

	if c00.A != 0 {
		t.Errorf(
			"Expected pixel (0,0) to be transparent, got alpha=%d",
			c00.A,
		)
	}

	// Check off-white background.
	c10 := color.RGBAModel.Convert(
		outImg.At(1, 0),
	).(color.RGBA)

	if c10.A != 0 {
		t.Errorf(
			"Expected pixel (1,0) to be transparent, got alpha=%d",
			c10.A,
		)
	}

	// Check foreground.
	c01 := color.RGBAModel.Convert(
		outImg.At(0, 1),
	).(color.RGBA)

	if c01.R != 255 ||
		c01.G != 0 ||
		c01.B != 0 ||
		c01.A != 255 {

		t.Errorf(
			"Expected foreground pixel to remain RGBA(255,0,0,255), got RGBA(%d,%d,%d,%d)",
			c01.R,
			c01.G,
			c01.B,
			c01.A,
		)
	}

	// Check bottom-right background.
	c11 := color.RGBAModel.Convert(
		outImg.At(1, 1),
	).(color.RGBA)

	if c11.A != 0 {
		t.Errorf(
			"Expected pixel (1,1) to be transparent, got alpha=%d",
			c11.A,
		)
	}
}

func TestConvertFolderFromPngToJpeg(t *testing.T) {
	tempDir := t.TempDir()

	inputDir := filepath.Join(tempDir, "input")
	outputDir := filepath.Join(tempDir, "output")

	if err := os.MkdirAll(inputDir, 0755); err != nil {
		t.Fatalf("Failed to create input directory: %v", err)
	}

	// Create PNG files.
	createDummyPNG(
		t,
		filepath.Join(inputDir, "image1.png"),
	)

	createDummyPNG(
		t,
		filepath.Join(inputDir, "image2.png"),
	)

	// Create unsupported file.
	txtPath := filepath.Join(inputDir, "ignore.txt")

	if err := os.WriteFile(
		txtPath,
		[]byte("this should be ignored"),
		0644,
	); err != nil {
		t.Fatalf("Failed to create text file: %v", err)
	}

	// Run conversion.
	err := ConvertFolderFromPngToJpeg(
		inputDir,
		outputDir,
		90,
	)

	if err != nil {
		t.Fatalf(
			"ConvertFolderFromPngToJpeg failed: %v",
			err,
		)
	}

	// Check image1.jpg.
	verifyJPEG(t, filepath.Join(
		outputDir,
		"image1.jpg",
	))

	// Check image2.jpg.
	verifyJPEG(t, filepath.Join(
		outputDir,
		"image2.jpg",
	))

	// Make sure unsupported file wasn't converted.
	ignoredPath := filepath.Join(
		outputDir,
		"ignore.jpg",
	)

	if _, err := os.Stat(ignoredPath); !os.IsNotExist(err) {
		t.Errorf(
			"Expected unsupported file to be ignored",
		)
	}
}

func TestConvertFolderFromJpegToPng(t *testing.T) {
	tempDir := t.TempDir()

	inputDir := filepath.Join(tempDir, "input")
	outputDir := filepath.Join(tempDir, "output")

	if err := os.MkdirAll(inputDir, 0755); err != nil {
		t.Fatalf("Failed to create input directory: %v", err)
	}

	// Create JPG file.
	createDummyJPEG(
		t,
		filepath.Join(inputDir, "image1.jpg"),
	)

	// Create JPEG file.
	createDummyJPEG(
		t,
		filepath.Join(inputDir, "image2.jpeg"),
	)

	// Create unsupported file.
	txtPath := filepath.Join(inputDir, "ignore.txt")

	if err := os.WriteFile(
		txtPath,
		[]byte("this should be ignored"),
		0644,
	); err != nil {
		t.Fatalf("Failed to create text file: %v", err)
	}

	// Run conversion.
	err := ConvertFolderFromJpegToPng(
		inputDir,
		outputDir,
	)

	if err != nil {
		t.Fatalf(
			"ConvertFolderFromJpegToPng failed: %v",
			err,
		)
	}

	// Check image1.png.
	verifyPNG(t, filepath.Join(
		outputDir,
		"image1.png",
	))

	// Check image2.png.
	verifyPNG(t, filepath.Join(
		outputDir,
		"image2.png",
	))

	// Make sure unsupported file wasn't converted.
	ignoredPath := filepath.Join(
		outputDir,
		"ignore.png",
	)

	if _, err := os.Stat(ignoredPath); !os.IsNotExist(err) {
		t.Errorf(
			"Expected unsupported file to be ignored",
		)
	}
}

func createDummyPNG(t *testing.T, path string) {
	t.Helper()

	img := image.NewRGBA(
		image.Rect(0, 0, 1, 1),
	)

	img.Set(
		0,
		0,
		color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		},
	)

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf(
			"Failed to create dummy PNG: %v",
			err,
		)
	}

	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		t.Fatalf(
			"Failed to encode dummy PNG: %v",
			err,
		)
	}
}

func createDummyJPEG(t *testing.T, path string) {
	t.Helper()

	img := image.NewRGBA(
		image.Rect(0, 0, 1, 1),
	)

	img.Set(
		0,
		0,
		color.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 255,
		},
	)

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf(
			"Failed to create dummy JPEG: %v",
			err,
		)
	}

	defer file.Close()

	err = jpeg.Encode(
		file,
		img,
		&jpeg.Options{
			Quality: 90,
		},
	)

	if err != nil {
		t.Fatalf(
			"Failed to encode dummy JPEG: %v",
			err,
		)
	}
}

func verifyJPEG(t *testing.T, path string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf(
			"Expected JPEG file was not created at %s: %v",
			path,
			err,
		)
	}

	if info.Size() == 0 {
		t.Fatalf(
			"JPEG file is empty: %s",
			path,
		)
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf(
			"Failed to open JPEG: %v",
			err,
		)
	}

	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		t.Fatalf(
			"Output is not a valid JPEG: %v",
			err,
		)
	}

	if img.Bounds().Dx() != 1 ||
		img.Bounds().Dy() != 1 {

		t.Errorf(
			"Expected JPEG size 1x1, got %dx%d",
			img.Bounds().Dx(),
			img.Bounds().Dy(),
		)
	}
}

func verifyPNG(t *testing.T, path string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf(
			"Expected PNG file was not created at %s: %v",
			path,
			err,
		)
	}

	if info.Size() == 0 {
		t.Fatalf(
			"PNG file is empty: %s",
			path,
		)
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf(
			"Failed to open PNG: %v",
			err,
		)
	}

	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatalf(
			"Output is not a valid PNG: %v",
			err,
		)
	}

	if img.Bounds().Dx() != 1 ||
		img.Bounds().Dy() != 1 {

		t.Errorf(
			"Expected PNG size 1x1, got %dx%d",
			img.Bounds().Dx(),
			img.Bounds().Dy(),
		)
	}
}