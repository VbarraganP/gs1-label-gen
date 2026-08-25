package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
)

func main() {
	// 1. Read the input file
	fileName := "texto.txt"
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading the file '%s': %v\n", fileName, err)
	}

	// 2. Parse the lines
	// Standardize line endings just in case (Windows/Linux)
	rawText := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(rawText), "\n")
	
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	if len(lines) < 3 {
		log.Fatalf("Error: The file '%s' must contain at least 3 lines (GTIN, Serial, Date).\n", fileName)
	}

	// Line 1: GTIN (AI 01 - must be 14 digits)
	gtin := lines[0]
	if len(gtin) < 14 {
		// Pad with leading zeros to meet the 14-digit GS1 requirement
		gtin = strings.Repeat("0", 14-len(gtin)) + gtin
	} else if len(gtin) > 14 {
		log.Fatalf("Error: GTIN cannot exceed 14 digits.")
	}

	// Line 2: Serial (AI 21)
	serial := lines[1]
	
	// Line 3: Expiration Date (AI 17 - YYMMDD)
	expDate := lines[2]
	if len(expDate) != 6 {
		log.Fatalf("Error: Expiration date must be exactly 6 digits (YYMMDD).")
	}

	// 3. Construct the GS1 String
	// We place fixed-length AIs first (01 and 17) and the variable-length AI last (21)
	// This strategic ordering avoids the need to insert a <GS> (Group Separator) character.
	// Format: FNC1 + 01 + [GTIN] + 17 + [DATE] + 21 + [SERIAL]
	gs1Data := string([]byte{datamatrix.FNC1}) + "01" + gtin + "17" + expDate + "21" + serial

	// 4. Generate the GS1 DataMatrix
	code, err := datamatrix.Encode(gs1Data)
	if err != nil {
		log.Fatalf("Error generating DataMatrix: %v\n", err)
	}

	// Scale the barcode to a readable physical size (200x200 pixels)
	// Because DataMatrix outputs perfect squares, this scaling maintains maximum sharpness.
	scaledCode, err := barcode.Scale(code, 200, 200)
	if err != nil {
		log.Fatalf("Error scaling DataMatrix: %v\n", err)
	}

	// 5. Add Quiet Zone (White Border)
	// DataMatrix requires a white border to be readable by scanners.
	bounds := scaledCode.Bounds()
	padding := 20 // 20 pixels white border on all sides
	paddedImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx()+(padding*2), bounds.Dy()+(padding*2)))
	
	// Fill with white
	for x := 0; x < paddedImage.Bounds().Dx(); x++ {
		for y := 0; y < paddedImage.Bounds().Dy(); y++ {
			paddedImage.Set(x, y, image.White)
		}
	}
	
	// Draw the barcode in the center
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			paddedImage.Set(x+padding, y+padding, scaledCode.At(x, y))
		}
	}

	// 6. Save to file
	qrFileName := "datamatrix.png"
	file, err := os.Create(qrFileName)
	if err != nil {
		log.Fatalf("Error creating file: %v\n", err)
	}
	defer file.Close()

	err = png.Encode(file, paddedImage)
	if err != nil {
		log.Fatalf("Error encoding PNG: %v\n", err)
	}

	// Cleanup old qrcode.png if it exists to avoid confusion
	os.Remove("qrcode.png")

	fmt.Printf("GS1 DataMatrix successfully generated! Saved as '%s'\n", qrFileName)
	fmt.Printf("Encoded Data (without FNC1 prefix): (01)%s (17)%s (21)%s\n", gtin, expDate, serial)
}
