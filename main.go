package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

func main() {
	// 1. Read the input file
	fileName := "texto.txt"
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading the file '%s': %v\n", fileName, err)
	}

	// Clean up leading/trailing whitespaces and newlines
	textData := strings.TrimSpace(string(content))

	if textData == "" {
		log.Fatalf("The file '%s' is empty.\n", fileName)
	}

	// 2. Generate the QR code
	// We use qrcode.Low for lower error correction, resulting in a code with fewer modules.
	// This makes it simpler and easier to scan at very small sizes.
	qrFileName := "qrcode.png"
	qr, err := qrcode.New(textData, qrcode.Low)
	if err != nil {
		log.Fatalf("Error preparing the QR code: %v\n", err)
	}

	// Remove the default white border (quiet zone) to maximize physical space usage
	qr.DisableBorder = true

	// Write the file. Using a negative value (e.g. -5) creates exact 5x5 pixel blocks.
	// This ensures maximum sharpness without blurriness, ideal for scaling down to a tiny physical size.
	err = qr.WriteFile(-5, qrFileName)
	if err != nil {
		log.Fatalf("Error generating the QR code: %v\n", err)
	}

	fmt.Printf("QR code successfully generated! Saved as '%s'\n", qrFileName)
	fmt.Printf("QR Content:\n%s\n", textData)
}
