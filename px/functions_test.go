package px

import (
	"fmt"
	"math"
	"testing"
)

func TestHslToRgb(t *testing.T) {

	if "000000" != hslToRgbStr(0,0,0) { t.Errorf("expected %v but got %v", "000000", hslToRgbStr(0,0,0)) }
	if "ffffff" != hslToRgbStr(0,0,1) { t.Errorf("expected %v but got %v", "ffffff", hslToRgbStr(0,0,1)) }

	if "000000" != hslToRgbStr(.5,0,0) { t.Errorf("expected %v but got %v", "000000", hslToRgbStr(.5,0,0)) }
	if "ffffff" != hslToRgbStr(.5,0,1) { t.Errorf("expected %v but got %v", "ffffff", hslToRgbStr(.5,0,1)) }

	if "000000" != hslToRgbStr(1,0,0) { t.Errorf("expected %v but got %v", "000000", hslToRgbStr(1,0,0)) }
	if "ffffff" != hslToRgbStr(1,0,1) { t.Errorf("expected %v but got %v", "ffffff", hslToRgbStr(1,0,1)) }

	if "ff0000" != hslToRgbStr(    0,1,.5) { t.Errorf("expected %v but got %v", "ff0000", hslToRgbStr(    0,1,.5)) } // red
	if "00ff00" != hslToRgbStr(1.0/3,1,.5) { t.Errorf("expected %v but got %v", "00ff00", hslToRgbStr(1.0/3,1,.5)) } // green
	if "0000ff" != hslToRgbStr(2.0/3,1,.5) { t.Errorf("expected %v but got %v", "0000ff", hslToRgbStr(2.0/3,1,.5)) } // blue

	// some online utitlities say bf4040 and b81414, other say different things... WFT?
	if "bf3f3f" != hslToRgbStr(0,.5,.5) { t.Errorf("expected %v but got %v", "bf3f3f", hslToRgbStr(0,.5,.5)) }
	// if "bf4040" != hslToRgbStr(0,.5,.5) { t.Errorf("expected %v but got %v", "bf4040", hslToRgbStr(0,.5,.5)) }

	if "b71414" != hslToRgbStr(0,.8,.4) { t.Errorf("expected %v but got %v", "b71414", hslToRgbStr(0,.8,.4)) }
	//if "b81414" != hslToRgbStr(0,.8,.4) { t.Errorf("expected %v but got %v", "b81414", hslToRgbStr(0,.8,.4)) }

}

func TestRgbToHsl(t *testing.T) {

	// Note hex codes below are HSL, _not_ RGB!!!

	if "000000" != rgbToHslStr(0,0,0) { t.Errorf("expected %v but got %v", "000000", rgbToHslStr(0,0,0)) }
	if "0000ff" != rgbToHslStr(1,1,1) { t.Errorf("expected %v but got %v", "0000ff", rgbToHslStr(1,1,1)) }

	if "00ff7f" != rgbToHslStr(1,0,0) { t.Errorf("expected %v but got %v", "00ffff", rgbToHslStr(1,0,0)) } // red
	if "55ff7f" != rgbToHslStr(0,1,0) { t.Errorf("expected %v but got %v", "55ff7f", rgbToHslStr(0,1,0)) } // green
	if "aaff7f" != rgbToHslStr(0,0,1) { t.Errorf("expected %v but got %v", "aaff7f", rgbToHslStr(0,0,1)) } // blue

}

func TestRgbHslRgb(t *testing.T) {
	cases := [][3]float64{
		{ .6, .3, .2 },
		{ .6, .3, .2 },
		{ .1,  0,  1 },
	}

	const TOLERANCE float64 = .25/255

	for _,RGB := range cases {
		hsl := rgbToHsl(RGB[0], RGB[1], RGB[2])
		rgb := hslToRgb(hsl[0], hsl[1], hsl[2])
		for i := range []int{0,1,2} {
			if math.Abs(RGB[i] - rgb[i]) > TOLERANCE {
				t.Errorf("%v: expected %v to be something like %v, but got %v", RGB, "RGB"[i:i+1], RGB[i], rgb[i])
			}
		}
	}
}

// --- here be utilities -----------------------------------------------------------------------------------------------

func hslToRgbStr(h float64, s float64, l float64) (string) {
	rgb := hslToRgb(h,s,l)
	b := func(f float64) (uint8) { return uint8(f * 255) }
	return fmt.Sprintf("%02x%02x%02x", b(rgb[0]), b(rgb[1]), b(rgb[2]))
}

func rgbToHslStr(r float64, g float64, b float64) (string) {
	hsl := rgbToHsl(r,g,b)
	bf := func(f float64) (uint8) { return uint8(f * 255) }
	return fmt.Sprintf("%02x%02x%02x", bf(hsl[0]), bf(hsl[1]), bf(hsl[2]))
}
