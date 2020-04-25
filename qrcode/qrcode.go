//  https://github.com/mdp/qrterminal
package qrcode

import (
	"io"
	"strings"

	"rsc.io/qr"
)

const WHITE = "\033[47m  \033[0m"
const BLACK = "\033[46;37m  \033[0m"

// Use ascii blocks to form the QR Code
const BLACK_WHITE = "▄"
const BLACK_BLACK = " "
const WHITE_BLACK = "▀"
const WHITE_WHITE = "█"

// Level - the QR Code's redundancy level
const H = qr.H
const M = qr.M
const L = qr.L

// default is 4-pixel-wide white quiet zone
const QUIET_ZONE = 1

//Config for generating a barcode
type Config struct {
	Level          qr.Level
	Writer         io.Writer
	HalfBlocks     bool
	BlackChar      string
	BlackWhiteChar string
	WhiteChar      string
	WhiteBlackChar string
	QuietZone      int
}

// Generate a QR Code and write it out to io.Writer
func Generate(text string, l qr.Level, w io.Writer) {
	config := Config{
		Level:      l,
		Writer:     w,
		BlackChar:  BLACK,
		WhiteChar:  WHITE,
		QuietZone:  QUIET_ZONE,
		HalfBlocks: false,
	}
	GenerateWithConfig(text, config)

}

// GenerateWithConfig expects a string to encode and a config
func GenerateWithConfig(text string, config Config) {
	if config.QuietZone < 1 {
		config.QuietZone = 1 // at least 1-pixel-wide white quiet zone
	}
	w := config.Writer
	code, _ := qr.Encode(text, config.Level)
	config.writeFullBlocks(w, code)

}

func stringRepeat(s string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(s, count)
}

func (c *Config) writeFullBlocks(w io.Writer, code *qr.Code) {
	white := c.WhiteChar
	black := c.BlackChar

	// Frame the barcode in a 1 pixel border
	w.Write([]byte(stringRepeat(stringRepeat(white,
		code.Size+c.QuietZone*2)+"\n", c.QuietZone))) // top border
	for i := 0; i <= code.Size; i++ {
		w.Write([]byte(stringRepeat(white, c.QuietZone))) // left border
		for j := 0; j <= code.Size; j++ {
			if code.Black(j, i) {
				w.Write([]byte(black))
			} else {
				w.Write([]byte(white))
			}
		}
		w.Write([]byte(stringRepeat(white, c.QuietZone-1) + "\n")) // right border
	}
	w.Write([]byte(stringRepeat(stringRepeat(white,
		code.Size+c.QuietZone*2)+"\n", c.QuietZone-1))) // bottom border
}
