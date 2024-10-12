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

func addAlignmentPattern(qrcode *qrcode) {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 || j == 0 || i == 4 || j == 4 {
				qrcode.data[len(qrcode.data)-9+i][len(qrcode.data)-9+j] = black
			} else {
				qrcode.data[len(qrcode.data)-9+i][len(qrcode.data)-9+j] = white
			}
		}
	}
	qrcode.data[len(qrcode.data)-7][len(qrcode.data)-7] = black
}

func addTimingStrips(qrcode *qrcode) {
	for i := 7; i < len(qrcode.data)-7; i++ {
		if i%2 == 1 {
			qrcode.data[6][i] = white
			qrcode.data[i][6] = white
		} else {
			qrcode.data[6][i] = black
			qrcode.data[i][6] = black
		}
	}
}
func newQrCode(len int) qrcode {
	qrCodeSize := 25
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

	addAlignmentPattern(&qrcode)

	addTimingStrips(&qrcode)

	return qrcode

}

func renderQRCode(qr qrcode) {
	for i := 0; i < len(qr.data)+2; i++ {
		fmt.Printf(whiteColor + "██" + defaultColor)
	}
	fmt.Println()
	for _, row := range qr.data {
		fmt.Printf(whiteColor + "██" + defaultColor)
		for _, col := range row {
			if col == black {
				fmt.Printf(blackColor + "██" + defaultColor)
			} else {
				fmt.Printf(whiteColor + "██" + defaultColor)
			}
		}
		fmt.Printf(whiteColor + "██" + defaultColor)
		fmt.Println()
	}
	for i := 0; i < len(qr.data)+2; i++ {
		fmt.Printf(whiteColor + "██" + defaultColor)

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
