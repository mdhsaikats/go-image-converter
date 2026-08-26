# imgconv

[![Go Reference](https://pkg.go.dev/badge/github.com/mdhsaikats/go-image-converter.svg)](https://pkg.go.dev/github.com/mdhsaikats/go-image-converter)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdhsaikats/go-image-converter)](https://goreportcard.com/report/github.com/mdhsaikats/go-image-converter)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`imgconv` is a lightweight, zero-dependency Go package for image format conversion (PNG <-> JPEG), bulk folder batch processing, and tolerance-based background removal.

---

## Features

- **Single Image Conversion**: Fast conversion between PNG and JPEG formats.
- **Bulk Folder Conversion**: Batch process entire directories of images (`PNG → JPEG` and `JPEG → PNG`) with automatic directory creation and non-image file filtering.
- **Background Removal**: Automatically sample top-left background color signatures and convert matching background pixels into transparent alpha.
- **Configurable Quality**: Custom JPEG compression quality scale (1-100).
- **Zero External Dependencies**: Implemented entirely with Go standard library packages (`image/jpeg`, `image/png`, `image/color`, `os`).

---

## Installation

Install the package using standard `go get`:

```bash
go get github.com/mdhsaikats/go-image-converter
```

---

## User Guide & Quick Start

### 1. Bulk Folder Conversions

#### Bulk PNG to JPEG (`ConvertFolderFromPngToJpeg`)

Batch converts all `.png` files inside `inputFolder` and outputs `.jpg` files into `outputFolder`:

```go
package main

import (
	"fmt"
	"log"

	imgconv "github.com/mdhsaikats/go-image-converter"
)

func main() {
	inputFolder := "./images/png_folder"
	outputFolder := "./images/jpeg_output"
	quality := 85 // JPEG compression quality (1-100)

	err := imgconv.ConvertFolderFromPngToJpeg(inputFolder, outputFolder, quality)
	if err != nil {
		log.Fatalf("Bulk PNG to JPEG conversion failed: %v", err)
	}

	fmt.Println("Bulk PNG to JPEG conversion completed successfully!")
}
```

#### Bulk JPEG to PNG (`ConvertFolderFromJpegToPng`)

Batch converts all `.jpg` and `.jpeg` files inside `inputFolder` and outputs `.png` files into `outputFolder`:

```go
package main

import (
	"fmt"
	"log"

	imgconv "github.com/mdhsaikats/go-image-converter"
)

func main() {
	inputFolder := "./images/jpeg_folder"
	outputFolder := "./images/png_output"

	err := imgconv.ConvertFolderFromJpegToPng(inputFolder, outputFolder)
	if err != nil {
		log.Fatalf("Bulk JPEG to PNG conversion failed: %v", err)
	}

	fmt.Println("Bulk JPEG to PNG conversion completed successfully!")
}
```

---

### 2. Single Image Conversions

#### PNG to JPEG (`PngToJpeg`)

```go
package main

import (
	"fmt"
	"log"

	imgconv "github.com/mdhsaikats/go-image-converter"
)

func main() {
	err := imgconv.PngToJpeg("input.png", "output.jpg", 90)
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	}

	fmt.Println("Converted single PNG to JPEG!")
}
```

#### JPEG to PNG (`JpegToPng`)

```go
package main

import (
	"fmt"
	"log"

	imgconv "github.com/mdhsaikats/go-image-converter"
)

func main() {
	err := imgconv.JpegToPng("input.jpg", "output.png")
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	}

	fmt.Println("Converted single JPEG to PNG!")
}
```

---

### 3. Background Removal (`RemoveBackground`)

```go
package main

import (
	"fmt"
	"log"

	imgconv "github.com/mdhsaikats/go-image-converter"
)

func main() {
	inputPath := "photo.png"
	outputPath := "photo_transparent.png"
	tolerance := uint8(15) // Color difference threshold (0 to 255)

	err := imgconv.RemoveBackground(inputPath, outputPath, tolerance)
	if err != nil {
		log.Fatalf("Background removal failed: %v", err)
	}

	fmt.Println("Background removed successfully!")
}
```

---

## API Reference

### Bulk Functions

#### `func ConvertFolderFromPngToJpeg(inputFolder, outputFolder string, quality int) error`
Reads all `.png` files in `inputFolder`, converts them to JPEG format using the specified `quality` setting, and saves them to `outputFolder`. Non-PNG files and subdirectories are automatically skipped. The output directory is created automatically if it does not exist.

#### `func ConvertFolderFromJpegToPng(inputFolder, outputFolder string) error`
Reads all `.jpg` and `.jpeg` files in `inputFolder`, converts them to PNG format, and saves them to `outputFolder`. Non-JPEG files and subdirectories are automatically skipped. The output directory is created automatically if it does not exist.

---

### Single File Functions

#### `func PngToJpeg(inputPath, outputPath string, quality int) error`
Converts a single PNG file at `inputPath` to a JPEG file at `outputPath` with quality scale `1` (lowest) to `100` (highest).

#### `func JpegToPng(inputPath, outputPath string) error`
Converts a single JPEG file at `inputPath` to a PNG file at `outputPath`.

#### `func RemoveBackground(inputPath, outputPath string, tolerance uint8) error`
Samples the top-left pixel `(0,0)` of the input image as the background color reference and turns all pixels matching within `tolerance` into transparent alpha in the output PNG file.

---

## License

This project is licensed under the MIT License.
