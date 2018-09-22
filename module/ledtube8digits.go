package module

import "awesomeProject/chip"

type LedTube8Digits struct {
	chip.Chip74hc595
	NumberToShow int
}

var (
	div      = []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000}
	position = []uint8{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}
	font     = []uint8{
		0xc0, 0xf9, 0xa4, 0xb0, 0x99, 0x92, 0x82, 0xf8,
		0x80, 0x90, 0x88, 0x83, 0xa7, 0xa1, 0x86, 0x8e}
)

func (led LedTube8Digits) Show() {
	var digits [8]uint8
	digits = led.splitNumberToShow()
	for i := 0; i < 8; i++ {
		var data uint16
		data = uint16(font[digits[i]])
		data = (data << 8) | uint16(position[i])
		led.SendUint16(data)
	}

}

func (led LedTube8Digits) splitNumberToShow() [8]uint8 {
	var res [8]uint8
	for i := 0; i < 8; i++ {
		res[i] = uint8((led.NumberToShow / div[i]) % 10)
	}
	return res
}
