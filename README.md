# imgconv

[![Go Reference](https://pkg.go.dev/badge/github.com/mdhsaikats/go-image-converter.svg)](https://pkg.go.dev/github.com/mdhsaikats/go-image-converter)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdhsaikats/go-image-converter)](https://goreportcard.com/report/github.com/mdhsaikats/go-image-converter)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`imgconv` is a lightweight, zero-dependency Go package for fast image format conversion between PNG and JPEG formats.

---

## Features

- **PNG to JPEG Conversion**: Convert PNG images to JPEG format with customizable quality compression settings (1-100).
- **JPEG to PNG Conversion**: Convert JPEG images to lossless PNG format cleanly and efficiently.
- **Zero External Dependencies**: Uses Go standard library packages (`image/jpeg`, `image/png`).
- **Idiomatic Error Handling**: Returns clear, wrapped errors for file access and image decoding/encoding operations.

---

## Installation

Install the package using standard `go get`:

```bash
go get github.com/mdhsaikats/go-image-converter
```

---

## User Guide & Quick Start

### 1. Converting PNG to JPEG (`PngToJpeg`)

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

### 2. Converting JPEG to PNG (`JpegToPng`)

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

### `func PngToJpeg(inputPath, outputPath string, quality int) error`

Converts a PNG image located at `inputPath` to a JPEG file saved at `outputPath`.

#### Parameters

- **`inputPath`** `(string)`: Path to the source PNG file.
- **`outputPath`** `(string)`: Destination path for the generated JPEG file.
- **`quality`** `(int)`: Compression quality setting for the output JPEG file (`1` to `100`).
  - **`1 - 30`**: High compression, smaller file size, lower visual quality.
  - **`75 - 85`**: Recommended balance between file size and image clarity.
  - **`90 - 100`**: Maximum visual quality, larger file size.

#### Return Value

- Returns `nil` on success, or a descriptive `error` if any stage of file reading, decoding, or encoding fails.

---

### `func JpegToPng(inputPath, outputPath string) error`

Converts a JPEG image located at `inputPath` to a PNG file saved at `outputPath`.

#### Parameters

- **`inputPath`** `(string)`: Path to the source JPEG file.
- **`outputPath`** `(string)`: Destination path for the generated PNG file.

#### Return Value

- Returns `nil` on success, or a descriptive `error` if any stage of file reading, decoding, or encoding fails.

---

## License

This project is licensed under the MIT License.
