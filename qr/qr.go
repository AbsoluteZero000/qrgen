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
			qrcode.data[7][i] = white
			qrcode.data[i][7] = white
		} else {
			qrcode.data[7][i] = black
			qrcode.data[i][7] = black
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

func addEncoding(datatype string, qrcode *qrcode) {
	qrSize := len(qrcode.data)
	qrcode.data[qrSize-1][qrSize-1] = white
	qrcode.data[qrSize-1][qrSize-2] = black
	qrcode.data[qrSize-2][qrSize-1] = white
	qrcode.data[qrSize-2][qrSize-2] = white
}

func addLength(dataLength int, qrcode *qrcode) {
	qrSize := len(qrcode.data)
	counter := 7

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			if (dataLength & (1 << counter)) > 0 {
				qrcode.data[qrSize-3-i][qrSize-j-1] = black
			} else {
				qrcode.data[qrSize-3-i][qrSize-j-1] = white
			}
			counter--
		}
	}
}

func toBitString(data string) string {

	result := ""

	for i := 0; i < len(data); i++ {
		result += fmt.Sprintf("%08b", data[i])
	}

	return result
}

func addData(data string, qrcode *qrcode) {
	// bitArray := toBitString(data)
	// counter := 0
	// totalCells := len(qrcode.data) * len(qrcode.data)

	// for j := len(qrcode.data) - 1; j >= 0; j -= 2 {
	// 	for i := len(qrcode.data) - 1; i >= 0; i-- {
	// 		if qrcode.data[i][j] != uncolored {
	// 			continue
	// 		}
	// 		if counter < len(bitArray) {
	// 			if bitArray[counter] == '1' {
	// 				qrcode.data[i][j] = black
	// 			} else {
	// 				qrcode.data[i][j] = white
	// 			}
	// 		} else {
	// 			// If we've used all data bits, fill with alternating pattern
	// 			qrcode.data[i][j] = black
	// 			if counter%2 == 0 {
	// 				qrcode.data[i][j] = white
	// 			}
	// 		}
	// 		counter++
	// 	}
	// 	if j > 0 {
	// 		for i := len(qrcode.data) - 1; i >= 0; i-- {
	// 			if qrcode.data[i][j-1] != uncolored {
	// 				continue
	// 			}
	// 			if counter < len(bitArray) {
	// 				if bitArray[counter] == '1' {
	// 					qrcode.data[i][j-1] = black
	// 				} else {
	// 					qrcode.data[i][j-1] = white
	// 				}
	// 			} else {
	// 				// If we've used all data bits, fill with alternating pattern
	// 				qrcode.data[i][j-1] = black
	// 				if counter%2 == 0 {
	// 					qrcode.data[i][j-1] = white
	// 				}
	// 			}
	// 			counter++
	// 		}
	// 	}
	// 	if counter >= totalCells {
	// 		return
	// 	}
	// }
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

	addEncoding("binary", &qrcode)

	addLength(len(data), &qrcode)

	// addData(data, &qrcode)

	return qrcode
}

func RenderQRCode(qr qrcode) {
	const (
		blackColor   = "\033[38;2;0;0;0m"
		redColor     = "\033[38;2;255;0;0m"
		whiteColor   = "\033[97m"
		defaultColor = "\033[0m"
	)

	for i := 0; i < len(qr.data)+4; i++ {
		fmt.Printf(whiteColor + "██" + defaultColor)
	}

	fmt.Println()
	for i := 0; i < len(qr.data)+4; i++ {
		fmt.Printf(whiteColor + "██" + defaultColor)
	}

	fmt.Println()
	for _, row := range qr.data {
		fmt.Printf(whiteColor + "████" + defaultColor)
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
		fmt.Printf(whiteColor + "████" + defaultColor)
		fmt.Println()
	}

	for i := 0; i < len(qr.data)+4; i++ {
		fmt.Printf(whiteColor + "██" + defaultColor)
	}

	fmt.Println()
	for i := 0; i < len(qr.data)+4; i++ {
		fmt.Printf(whiteColor + "██" + defaultColor)
	}

}
