package main

import (
	"fmt"
)

// Function to render the QR code in the terminal
func renderQRCode(qr [][]int) {
	for _, row := range qr {
		for _, col := range row {
			if col == 1 {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func main() {
	qrCode := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}

	renderQRCode(qrCode)
}
