# GS1 DataMatrix Label Generator

A specialized Go application that generates industry-standard GS1 DataMatrix barcodes.

## Features
- Reads structured GS1 data (GTIN, Serial Number, and Expiration Date) from a `texto.txt` file.
- Automatically pads GTINs to 14 digits as required by GS1 specifications.
- Strategically orders Application Identifiers (AIs) to avoid needing Group Separator (`<GS>`) characters.
- Injects the mandatory `FNC1` character at the beginning of the barcode to ensure it is recognized by hardware scanners as a valid GS1 DataMatrix.
- Outputs a crisp, scalable `datamatrix.png` image.

## Requirements
- [Go](https://go.dev/) installed on your machine.

## Usage

### Option 1: Run directly
1. Ensure you are in the project directory.
2. Edit the `texto.txt` file in the root directory. You **must** provide exactly three lines of data in this order:
   - **Line 1:** GTIN (Product Code) - *Will be auto-padded to 14 digits if shorter.*
   - **Line 2:** Serial Number (AI 21)
   - **Line 3:** Expiration Date (AI 17) - *Must be exactly 6 digits in YYMMDD format.*
3. Run the application:
   ```bash
   go run main.go
   ```
4. The generated GS1 DataMatrix will be saved as `datamatrix.png` in the same directory.

### Option 2: Build and run the binary
If you prefer to compile the application into a standalone executable:
1. Build the binary:
   ```bash
   go build -o gs1-label-gen
   ```
2. Run the compiled binary:
   - **Linux / macOS:** `./gs1-label-gen`
   - **Windows:** `gs1-label-gen.exe`

## Technical Details (AIs)
The application automatically encodes the data using the following structure:
`[FNC1] (01) GTIN (17) YYMMDD (21) SERIAL`
