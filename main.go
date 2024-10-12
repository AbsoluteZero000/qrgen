package main

import (
	"fmt"
	"github.com/absolutezero000/qrcode-generator-reader/qr"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("usage: qrcode <data>")
		os.Exit(1)
	}
	data := os.Args[1]

	qrCode := qr.NewQrCode(data)
	qr.RenderQRCode(qrCode)
}
