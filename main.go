package main

import (
	"fmt"
)

type color int

const (
	uncolored color = iota
	white
	black
)

const (
	blackColor   = "\033[38;2;0;0;0m"
	whiteColor   = "\033[97m"
	defaultColor = "\033[0m"
)

func renderQRCode(qr [][]color) {
	for _, row := range qr {
		for _, col := range row {
			if col == white {
				fmt.Printf(whiteColor + "██" + defaultColor)
			} else {
				fmt.Printf(blackColor + "██" + defaultColor)
			}
		}
		fmt.Println()
	}
}

func main() {
	qrCode := [][]color{
		{white, white, white, white, white, white, white},
		{white, white, black, white, black, white, white},
		{white, white, white, white, white, white, white},
		{white, black, white, white, white, black, white},
		{white, black, white, white, white, black, white},
		{white, black, black, black, black, black, white},
		{white, white, white, white, white, white, white},
	}

	renderQRCode(qrCode)
}
