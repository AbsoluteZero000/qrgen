package main

import (
	"fmt"
)

func renderQRCode(qr [][]bool) {
	for _, row := range qr {
		for _, col := range row {
			if col == true {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func main() {
	qrCode := [][]bool{
		{true, true, true, true, true, true, true},
		{true, true, false, true, false, true, true},
		{true, true, true, true, true, true, true},
		{true, false, true, true, true, false, true},
		{true, false, true, true, true, false, true},
		{true, false, false, false, false, false, true},
		{true, true, true, true, true, true, true},
	}

	renderQRCode(qrCode)
}
