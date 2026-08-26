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
	// 1. Create a temporary directory that Go will automatically clean up after the test
	tempDir := t.TempDir()
	inputPath := filepath.Join(tempDir, "test.png")
	outputPath := filepath.Join(tempDir, "test.jpg")

	// 2. Create a dummy PNG file to use as input
	createDummyPNG(t, inputPath)

	// 3. Test your package's conversion function
	err := PngToJpeg(inputPath, outputPath, 90)
	if err != nil {
		t.Fatalf("PngToJpeg failed: %v", err)
	}

	// 4. Verify the output file was actually created and is not empty
	info, err := os.Stat(outputPath)
	if os.IsNotExist(err) {
		t.Fatalf("Expected output file was not created at %s", outputPath)
	}
	if info.Size() == 0 {
		t.Errorf("Output file is completely empty")
	}
}

// Helper function to create a simple 1x1 pixel PNG
func createDummyPNG(t *testing.T, path string) {
	t.Helper() // Tells Go this is a helper, so error lines point to the main test

	// Create a 1x1 red image in memory
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})

	// Create the physical file
	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("Failed to create dummy PNG: %v", err)
	}
	defer file.Close()

	// Write the PNG data to the file
	err = png.Encode(file, img)
	if err != nil {
		t.Fatalf("Failed to encode dummy PNG: %v", err)
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
	if os.IsNotExist(err) {
		t.Fatalf("Expected output file was not created at %s", outputPath)
	}
	if info.Size() == 0 {
		t.Errorf("Output file is completely empty")
	}
}

func TestRemoveBackground(t *testing.T) {
	tempDir := t.TempDir()
	inputPath := filepath.Join(tempDir, "bg_test.png")
	outputPath := filepath.Join(tempDir, "bg_removed.png")

	// Create a 2x2 image:
	// (0,0): White (255, 255, 255, 255) -> Background reference
	// (1,0): Off-white (250, 250, 250, 255) -> Within tolerance (diff 5 <= 10)
	// (0,1): Solid Red (255, 0, 0, 255) -> Foreground subject
	// (1,1): White (255, 255, 255, 255) -> Background
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{255, 255, 255, 255})
	img.Set(1, 0, color.RGBA{250, 250, 250, 255})
	img.Set(0, 1, color.RGBA{255, 0, 0, 255})
	img.Set(1, 1, color.RGBA{255, 255, 255, 255})

	file, err := os.Create(inputPath)
	if err != nil {
		t.Fatalf("Failed to create input PNG: %v", err)
	}
	if err := png.Encode(file, img); err != nil {
		file.Close()
		t.Fatalf("Failed to encode input PNG: %v", err)
	}
	file.Close()

	// Execute background removal with tolerance of 10
	if err := RemoveBackground(inputPath, outputPath, 10); err != nil {
		t.Fatalf("RemoveBackground failed: %v", err)
	}

	// Verify generated transparent PNG
	outFile, err := os.Open(outputPath)
	if err != nil {
		t.Fatalf("Failed to open output image: %v", err)
	}
	defer outFile.Close()

	outImg, err := png.Decode(outFile)
	if err != nil {
		t.Fatalf("Failed to decode output image: %v", err)
	}

	// Check top-left background pixel (0,0) is now transparent
	c00 := color.RGBAModel.Convert(outImg.At(0, 0)).(color.RGBA)
	if c00.A != 0 {
		t.Errorf("Expected pixel (0,0) to be transparent (alpha=0), got alpha=%d", c00.A)
	}

	// Check top-right off-white pixel (1,0) is also transparent
	c10 := color.RGBAModel.Convert(outImg.At(1, 0)).(color.RGBA)
	if c10.A != 0 {
		t.Errorf("Expected pixel (1,0) to be transparent (alpha=0), got alpha=%d", c10.A)
	}

	// Check bottom-left foreground red pixel (0,1) retained original color and full opacity
	c01 := color.RGBAModel.Convert(outImg.At(0, 1)).(color.RGBA)
	if c01.R != 255 || c01.G != 0 || c01.B != 0 || c01.A != 255 {
		t.Errorf("Expected foreground pixel (0,1) to remain RGBA(255,0,0,255), got RGBA(%d,%d,%d,%d)", c01.R, c01.G, c01.B, c01.A)
	}
}

func createDummyJPEG(t *testing.T, path string) {
	t.Helper()

	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{0, 255, 0, 255})

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("Failed to create dummy JPEG: %v", err)
	}
	defer file.Close()

	var opt jpeg.Options
	opt.Quality = 90
	err = jpeg.Encode(file, img, &opt)
	if err != nil {
		t.Fatalf("Failed to encode dummy JPEG: %v", err)
	}
}
