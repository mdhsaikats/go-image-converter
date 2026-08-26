# imgconv

[![Go Reference](https://pkg.go.dev/badge/github.com/mdhsaikats/go-image-converter.svg)](https://pkg.go.dev/github.com/mdhsaikats/go-image-converter)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdhsaikats/go-image-converter)](https://goreportcard.com/report/github.com/mdhsaikats/go-image-converter)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`imgconv` is a lightweight, zero-dependency Go package for image format conversion (PNG <-> JPEG) and simple background removal.

---

## Features

- **PNG to JPEG Conversion**: Convert PNG images to JPEG with configurable quality compression (1-100).
- **JPEG to PNG Conversion**: Convert JPEG images to lossless PNG format.
- **Background Removal**: Automatically detect top-left background color signatures and make matching background pixels transparent in PNG outputs with customizable color tolerance.
- **Zero External Dependencies**: Pure Go standard library implementation (`image/jpeg`, `image/png`, `image/color`).

---

## Installation

Install the package using standard `go get`:

```bash
go get github.com/mdhsaikats/go-image-converter
```

---

## User Guide & Quick Start

### 1. Removing Image Background (`RemoveBackground`)

Detects the background color starting from the top-left corner `(0,0)` and converts all matching pixels within the specified color tolerance to transparent alpha.

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
	tolerance := uint8(15) // Color difference threshold (0 = exact match, higher values match broader color variations)

	err := imgconv.RemoveBackground(inputPath, outputPath, tolerance)
	if err != nil {
		log.Fatalf("Failed to remove background: %v", err)
	}

	fmt.Println("Background removed successfully!")
}
```

### 2. Converting PNG to JPEG (`PngToJpeg`)

```go
package main

import (
	"fmt"
	"log"

	imgconv "github.com/mdhsaikats/go-image-converter"
)

func main() {
	inputPath := "sample.png"
	outputPath := "output.jpg"
	quality := 85 // Quality scale: 1 (lowest quality, highest compression) to 100 (highest quality)

	err := imgconv.PngToJpeg(inputPath, outputPath, quality)
	if err != nil {
		log.Fatalf("Failed to convert PNG to JPEG: %v", err)
	}

	fmt.Println("PNG converted to JPEG successfully!")
}
```

### 3. Converting JPEG to PNG (`JpegToPng`)

```go
package main

import (
	"fmt"
	"log"

	imgconv "github.com/mdhsaikats/go-image-converter"
)

func main() {
	inputPath := "sample.jpg"
	outputPath := "output.png"

	err := imgconv.JpegToPng(inputPath, outputPath)
	if err != nil {
		log.Fatalf("Failed to convert JPEG to PNG: %v", err)
	}

	fmt.Println("JPEG converted to PNG successfully!")
}
```

---

## API Reference

### `func RemoveBackground(inputPath, outputPath string, tolerance uint8) error`

Inspects the source image at `inputPath`, samples the top-left pixel `(0,0)` as the reference background color, converts all pixels within the RGB `tolerance` threshold to transparent alpha `RGBA(0, 0, 0, 0)`, and saves the resulting image to `outputPath` as a PNG.

#### Parameters

- **`inputPath`** `(string)`: Path to the input image file (PNG or JPEG).
- **`outputPath`** `(string)`: Destination path for the output PNG file (must support alpha transparency).
- **`tolerance`** `(uint8)`: Maximum allowed difference per RGB channel (0 to 255) when comparing pixels against the sampled background color.
  - **`0`**: Only pixels matching the exact background color are made transparent.
  - **`10 - 30`**: Recommended for solid backgrounds with slight compression artifacts or lighting variations.

#### Return Value

- Returns `nil` on success, or a descriptive `error` if image reading, decoding, or PNG encoding fails.

---

### `func PngToJpeg(inputPath, outputPath string, quality int) error`

Converts a PNG image located at `inputPath` to a JPEG file saved at `outputPath`.

#### Parameters

- **`inputPath`** `(string)`: Path to the source PNG file.
- **`outputPath`** `(string)`: Destination path for the generated JPEG file.
- **`quality`** `(int)`: Compression quality setting for the output JPEG file (`1` to `100`).

#### Return Value

- Returns `nil` on success, or a descriptive `error` if any stage fails.

---

### `func JpegToPng(inputPath, outputPath string) error`

Converts a JPEG image located at `inputPath` to a PNG file saved at `outputPath`.

#### Parameters

- **`inputPath`** `(string)`: Path to the source JPEG file.
- **`outputPath`** `(string)`: Destination path for the generated PNG file.

#### Return Value

- Returns `nil` on success, or a descriptive `error` if any stage fails.

---

## License

This project is licensed under the MIT License.
