package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"time"
	_  "image/gif"
	_  "image/jpeg"
	_  "image/png"
	"github.com/giorgiga/palette/options"
	"github.com/giorgiga/palette/px"
)

func main() {

	opts   := options.CliOptions()

	if opts.NColors < 1 {
		if opts.Verbose { fmt.Printf("%v colors? Whatever.\n", opts.NColors) }
		os.Exit(0)
	}

	var dispersionF func (*px.PixelSet) (int,float64)
	switch opts.Dispersion {
		case options.DISPERSION_SPREAD:   dispersionF = (*px.PixelSet).MaxSpread
		case options.DISPERSION_VARIANCE: fallthrough
		default:                          dispersionF = (*px.PixelSet).MaxVariance
	}

	pixels := loadPixels(opts)
	pset   := makeSet(opts, pixels)

	palette  := []px.PixelSet{ pset }
	for len(palette) < opts.NColors {
		palette = cut(opts, palette, dispersionF)
	}

	printOut(opts, palette)
}

func loadPixels(opts options.Options) []px.Pixel {

	startTime := time.Now()

	var file *os.File

	if opts.File == "" {
		file = os.Stdin
	} else {
		f, err := os.Open(opts.File)
		defer file.Close()
		if (err != nil) {
			fmt.Fprintf(os.Stderr, "Could not read %s (%s)\n", opts.File, err.Error())
			os.Exit(1)
		}
		file = f
	}

	image, _, err := image.Decode(file)
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "Could not read %s (%s)\n", opts.File, err.Error())
		os.Exit(1)
	}

	bounds := image.Bounds()
	size   := uint64(bounds.Max.Y - bounds.Min.Y) * uint64(bounds.Max.X - bounds.Min.X)
	pixels := make([]px.Pixel, 0, size)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r,g,b,a := image.At(x,y).RGBA() // alpha-premultiplied uint32s within [0, 0xffff]
			switch a {
				case 0xffff:          // 100% opaque pixel: r,g,b are the actual values
				case 0x0000: continue // 100% transparent pixel: skip it
				default: { r = (r * 0xffff) / a; g = (g * 0xffff) / a; b = (b * 0xffff) / a } // undo alpha premultiplication
			}
			pixels = append(pixels, px.Pixel{ uint8(r >> 8), uint8(g >> 8), uint8(b >> 8) })
		}
	}

	if len(pixels) == 0 {
		fmt.Fprintf(os.Stderr, "Could not read %s (%s)\n", opts.File, "all pixels are transparent")
		os.Exit(1)
	}

	if opts.Debug {
		elapsed := time.Now().Sub(startTime)
		fmt.Printf("%.2fM pixels read in %dms.\n\n", float64(len(pixels)) / 1e6, elapsed/time.Millisecond)
	} else if opts.Verbose {
		fmt.Printf("%.2fM pixels read.\n\n", float64(len(pixels)) / 1e6)
	}

	return pixels
}

func makeSet(opts options.Options, pixels []px.Pixel) px.PixelSet {
	startTime := time.Now()
	pset := px.MakePixelSet(pixels)
	if opts.Debug {
		elapsed := time.Now().Sub(startTime)
		fmt.Printf("makeSet: done in %dms\n", elapsed/time.Millisecond)
	}
	return pset
}

func cut(opts options.Options, psets []px.PixelSet, dispersionF func (*px.PixelSet) (int,float64)) []px.PixelSet {
	startTime := time.Now()

	maxspread := struct { index int; axis int; value float64 } { 0, 0, math.Inf(-1) }
	for i,pset := range psets {
		axis,spread := dispersionF(&pset)
		if spread > maxspread.value {
			maxspread.index = i
			maxspread.axis  = axis
			maxspread.value = spread
		}
	}
	head := psets[:maxspread.index]
	_, left, right := psets[maxspread.index].MedianCut(maxspread.axis)
	tail := psets[maxspread.index+1:]
	psets = make([]px.PixelSet,0,len(psets)+1)
	psets = append(psets, head...)
	psets = append(psets, left, right)
	psets = append(psets, tail...)

	if opts.Debug {
		elapsed := time.Now().Sub(startTime)
		fmt.Printf("cut %d: split on channel %v in %dms\n", len(psets)-1, "RGB"[maxspread.axis:maxspread.axis+1], elapsed/time.Millisecond)
	}

	return psets
}

func printOut(opts options.Options, psets []px.PixelSet) {

	if opts.Debug { fmt.Printf("\n") }
	if opts.Verbose { fmt.Printf("Palette:\n") }
	for i, pset := range psets {
		px := pset.Mean()

		var color string
		switch opts.Format {
			case options.FORMAT_D:
				color = fmt.Sprintf("%d,%d,%d", px[0], px[1], px[2])
			case options.FORMAT_F:
				color = fmt.Sprintf("%.6f,%.6f,%.6f", float64(px[0])/255, float64(px[1])/255, float64(px[2])/255)
			case options.FORMAT_X:
				color = fmt.Sprintf("%x", px)
		}

		if opts.Colorise {
			fg := "\033[38;2;0;0;0m" // black
			if (px.HSL()[2] < .5) { fg = "\033[38;2;255;255;255m" } // white
			spacer := ""
			if opts.Verbose { spacer = " " }
			color = fmt.Sprintf("%s\033[48;2;%d;%d;%dm%s\033[0m", fg, px[0], px[1], px[2], spacer + color + spacer)
		}

		if opts.Verbose {
			l := len(fmt.Sprintf("%d", opts.NColors))
			fmt.Printf("%" + fmt.Sprintf("%d",l) + "v. %s\n", i+1, color)
		} else {
			fmt.Printf("%s\n", color)
		}
	}

}
