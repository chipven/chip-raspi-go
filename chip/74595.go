package chip

import "github.com/stianeikeland/go-rpio"

//noinspection GoNameStartsWithPackageName
type Chip74hc595 struct {
	DIO rpio.Pin
	RCK rpio.Pin
	SCK rpio.Pin
}

func (c595 Chip74hc595) SendUint16(data uint16) {
	c595.DIO.Output()
	c595.RCK.Output()
	c595.SCK.Output()
	for i := 0; i < 16; i++ {
		c595.DIO.Write(rpio.State(uint8(data >> 15)))
		data = data << 1
		c595.SCK.Low()
		c595.SCK.High()
	}
	c595.RCK.Low()
	c595.RCK.High()
}
