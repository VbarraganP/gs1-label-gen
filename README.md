# QR Code Generator

A simple and efficient Go application that generates a physical-space-optimized QR code from plain text.

## Features
- Reads input from a `texto.txt` file.
- Generates a QR code (`qrcode.png`) using `github.com/skip2/go-qrcode`.
- Uses `qrcode.Low` error correction to minimize visual density, making the code easily scannable at very small physical sizes.
- Removes the quiet zone (white border) to maximize the use of available physical space.
- Configured to output "pixel-perfect" blocks (e.g., 5x5 pixels per module) to prevent blurring when the image is scaled down for printing.

## Requirements
- [Go](https://go.dev/) installed on your machine.

## Usage

### Option 1: Run directly
1. Ensure you are in the project directory.
2. Create or edit the `texto.txt` file in the root directory and add the content you want to encode. Keep it as short as possible for smaller QR codes.
3. Run the application:
   ```bash
   go run main.go
   ```
4. The generated QR code will be saved as `qrcode.png` in the same directory.

### Option 2: Build and run the binary
If you prefer to compile the application into a standalone executable:
1. Build the binary:
   ```bash
   go build -o qr-generator
   ```
2. Run the compiled binary:
   - **Linux / macOS:** `./qr-generator`
   - **Windows:** `qr-generator.exe`

## Customizing the QR Size

To change the pixel density of the generated QR code, modify the following line in `main.go`:
```go
err = qr.WriteFile(-5, qrFileName)
```
- `-5`: Means each module (QR block) will be exactly 5x5 pixels.
- Use smaller negative numbers like `-1` or `-2` for a tiny image file output, or positive numbers (like `256`) to set a fixed total width/height in pixels.
