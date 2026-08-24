package imgconv

import (
	"image"
	"image/color"
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