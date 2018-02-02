package px

import (
	"github.com/giorgiga/sstats"
)

// adapted from the ABC in https://www.w3.org/TR/css-color-3/#hsl-color

func hslToRgb (h float64, s float64, l float64) [3]float64 {
	// HOW TO RETURN hsl.to.rgb(h, s, l):
	//     SELECT:
	//         l<=0.5: PUT l*(s+1) IN m2
	//         ELSE: PUT l+s-l*s IN m2
	var m2 float64; if l < .5 { m2 = l * (s+1) } else { m2 = l + s - l * s }
	//     PUT l*2-m2 IN m1
 	var m1 float64; m1 = l * 2 - m2
	//     PUT hue.to.rgb(m1, m2, h+1/3) IN r
	r := hueToRgb(m1, m2, h + 1.0/3)
	//     PUT hue.to.rgb(m1, m2, h    ) IN g
	g := hueToRgb(m1, m2, h)
	//     PUT hue.to.rgb(m1, m2, h-1/3) IN b
	b := hueToRgb(m1, m2, h - 1.0/3)
	//     RETURN (r, g, b)
	return [3]float64 { r, g, b }
}

func hueToRgb (m1 float64, m2 float64, h float64) float64 {
	hh := h
	// HOW TO RETURN hue.to.rgb(m1, m2, h):
	// IF h<0: PUT h+1 IN h
	if hh < 0 { hh += 1 }
	// IF h>1: PUT h-1 IN h
	if hh > 1 { hh -= 1 }
	// IF h*6<1: RETURN m1+(m2-m1)*h*6
	if hh * 6 < 1 { return m1 + (m2 - m1) * hh * 6 }
	// IF h*2<1: RETURN m2
	if hh * 2 < 1 { return m2 }
	// IF h*3<2: RETURN m1+(m2-m1)*(2/3-h)*6
	if hh * 3 < 2 { return m1 + (m2 - m1) * (2.0/3 - hh) * 6 }
	// RETURN m1
	return m1
}

func rgbToHsl (r float64, g float64, b float64) [3]float64 {
	rgb := [3]float64{ r , g, b }
	maxi,maxv := sstats.Max(rgb[:]...)
	_,   minv := sstats.Min(rgb[:]...)
	diff := maxv - minv

	if (diff == 0) { return [3]float64{ 0, 0, (minv + maxv) / 2 } }

	l := (minv + maxv) / 2
	s := diff; if l > .5 { s /= 2 - 2 * l } else { s /= 2 * l }
	h := (rgb[(maxi+1)%3] - rgb[(maxi+2)%3]) / diff
	switch maxi {
		case 0: if rgb[1] < rgb[2] { h += 6 }
		case 1: h += 2
		case 2: h += 4
	}
	h /= 6
	return [3]float64{ h, s, l }
}
