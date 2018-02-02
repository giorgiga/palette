package px

import (
	"fmt"
)

// Pixel represents a 24bit (1 byte per channel) RGB color value.
type Pixel [3]uint8

// MakePXHSL converts the given HSL values (range 0..1 for all components)
// to RGB and then creates a Pixel
func MakePXHSL (h float64, s float64, l float64) Pixel {
	rgb := hslToRgb(h,s,l)
	return Pixel { uint8(rgb[0] * 255), uint8(rgb[1] * 255), uint8(rgb[2] * 255) }
}

// Format
func (px Pixel) Format(f fmt.State, c rune) {
	switch(c) {
		case 'x': f.Write([]byte( fmt.Sprintf("%02x%02x%02x", px[0], px[1], px[2]) ))
		case 'X': f.Write([]byte( fmt.Sprintf("%02X%02X%02X", px[0], px[1], px[2]) ))
		case 'v':
			if f.Flag('#') {
				f.Write([]byte( fmt.Sprintf("px.Pixel { %d, %d, %d }", px[0], px[1], px[2]) ))
			} else {
				px.Format(f, 'x')
			}
		default: px.Format(f, 'x')
	}
}

func (px Pixel) HSL() ([3]float64) {
	return rgbToHsl( float64(px[0])/255,
	                 float64(px[1])/255,
	                 float64(px[2])/255 )
}
