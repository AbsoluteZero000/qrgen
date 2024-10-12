package main

import (
	"fmt"
	"os"
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

type qrcode struct {
	data [][]color
}

func addFinderPattern(qr *qrcode) {
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if i == 0 || i == 6 || j == 0 || j == 6 {
				qr.data[i][j] = black
				qr.data[len(qr.data)-7+i][j] = black
				qr.data[i][len(qr.data)-7+j] = black
			} else if i == 1 || j == 1 || i == 5 || j == 5 {
				qr.data[i][j] = white
				qr.data[len(qr.data)-7+i][j] = white
				qr.data[i][len(qr.data)-7+j] = white
			} else {
				qr.data[i][j] = black
				qr.data[len(qr.data)-7+i][j] = black
				qr.data[i][len(qr.data)-7+j] = black
			}
		}
	}
}
func newQrCode(len int) qrcode {
	qrCodeSize := 21
	qrcode := qrcode{
		data: make([][]color, qrCodeSize),
	}

	for i := 0; i < qrCodeSize; i++ {
		qrcode.data[i] = make([]color, qrCodeSize)
		for j := 0; j < len; j++ {
			qrcode.data[i][j] = uncolored
		}
	}

	addFinderPattern(&qrcode)

	return qrcode

}
func renderQRCode(qr qrcode) {
	for _, row := range qr.data {
		for _, col := range row {
			if col == black {
				fmt.Printf(blackColor + "██" + defaultColor)
			} else {
				fmt.Printf(whiteColor + "██" + defaultColor)
			}
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("usage: qrcode <data>")
		os.Exit(1)
	}
	data := os.Args[1]

	qrCode := newQrCode(len(data))
	renderQRCode(qrCode)
}
