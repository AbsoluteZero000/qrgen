package qr

import (
	"fmt"
)

type color int

const (
	uncolored color = iota
	white
	black
	red
)

const (
	blackColor   = "\033[38;2;0;0;0m"
	redColor     = "\033[38;2;255;0;0m"
	whiteColor   = "\033[97m"
	defaultColor = "\033[0m"
)

type qrcode struct {
	data [][]color
}

func addFinderPattern(qr *qrcode) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == 7 || j == 7 {
				qr.data[i][j] = white
				if i == 7 {
					qr.data[i][len(qr.data)-8+j] = white
					qr.data[len(qr.data)-8][j] = white
				} else {
					qr.data[i][len(qr.data)-8] = white
					qr.data[len(qr.data)-7+i][j] = white
				}
			} else if i == 0 || i == 6 || j == 0 || j == 6 {
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

func addHeaders(data string, qrcode *qrcode) {

	for i := 0; i < 9; i++ {

		if qrcode.data[8][i] == uncolored {
			qrcode.data[8][i] = red
		}

		if qrcode.data[i][8] == uncolored {
			qrcode.data[i][8] = red
		}
		if i != 0 {
			if qrcode.data[8][len(qrcode.data)-9+i] == uncolored {
				qrcode.data[8][len(qrcode.data)-9+i] = red
			}

			if qrcode.data[len(qrcode.data)-9+i][8] == uncolored {
				qrcode.data[len(qrcode.data)-9+i][8] = red
			}
		}
	}

}

func NewQrCode(data string) qrcode {
	qrCodeSize := 25
	qrcode := qrcode{
		data: make([][]color, qrCodeSize),
	}

	for i := 0; i < qrCodeSize; i++ {
		qrcode.data[i] = make([]color, qrCodeSize)
		for j := 0; j < qrCodeSize; j++ {
			qrcode.data[i][j] = uncolored
		}
	}

	addFinderPattern(&qrcode)

	addAlignmentPattern(&qrcode)

	addTimingStrips(&qrcode)

	addHeaders(data, &qrcode)

	// Add useless dot
	qrcode.data[len(qrcode.data)-8][8] = black

	return qrcode

}

func RenderQRCode(qr qrcode) {
	for i := 0; i < len(qr.data)+2; i++ {
		fmt.Printf(whiteColor + "██" + defaultColor)
	}
	fmt.Println()
	for _, row := range qr.data {
		fmt.Printf(whiteColor + "██" + defaultColor)
		for _, col := range row {
			if col == black {
				fmt.Printf(blackColor + "██" + defaultColor)
			} else if col == white {
				fmt.Printf(whiteColor + "██" + defaultColor)
			} else if col == red {
				fmt.Printf(redColor + "██" + defaultColor)
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf(whiteColor + "██" + defaultColor)
		fmt.Println()
	}
	for i := 0; i < len(qr.data)+2; i++ {
		fmt.Printf(whiteColor + "██" + defaultColor)

	}
}
