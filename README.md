# imgconv

[![Go Reference](https://pkg.go.dev/badge/github.com/mdhsaikats/go-image-converter.svg)](https://pkg.go.dev/github.com/mdhsaikats/go-image-converter)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdhsaikats/go-image-converter)](https://goreportcard.com/report/github.com/mdhsaikats/go-image-converter)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`imgconv` is a lightweight, zero-dependency Go package designed for fast and simple PNG to JPEG image conversion with configurable compression quality.

---

## Installation

Install the package using standard `go get`:

```bash
go get github.com/mdhsaikats/go-image-converter
```

---

## User Guide & Quick Start

### Basic Conversion

To convert a PNG image to JPEG format, import `github.com/mdhsaikats/go-image-converter` (or alias it as `imgconv`) and call `PngToJpeg`:

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
		log.Fatalf("Failed to convert image: %v", err)
	}

	fmt.Println("Image converted successfully!")
}
```

---

## API Reference

### `func PngToJpeg(inputPath, outputPath string, quality int) error`

Converts a PNG image located at `inputPath` to a JPEG file saved at `outputPath`.

#### Parameters

- **`inputPath`** `(string)`: The filesystem path to the source PNG file.
- **`outputPath`** `(string)`: The destination path where the generated JPEG file will be saved.
- **`quality`** `(int)`: Compression quality setting for the output JPEG file, ranging from `1` to `100`.
  - **`1 - 30`**: High compression, smaller file size, reduced visual quality.
  - **`75 - 85`**: Recommended balance between file size and visual fidelity.
  - **`90 - 100`**: Maximum visual quality, larger file size.

#### Return Value

- Returns `nil` on success.
- Returns a descriptive `error` if opening the input file, decoding PNG data, creating the output file, or encoding JPEG data fails.

---

## Error Handling

`PngToJpeg` returns wrapped errors providing context on what failed during execution. It is good practice to handle errors explicitly:

```go
if err := imgconv.PngToJpeg("input.png", "output.jpg", 80); err != nil {
	log.Printf("Conversion error: %v", err)
}
```

---

## License

This project is licensed under the MIT License.
