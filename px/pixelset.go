package px

import (
	"math"
	"github.com/giorgiga/sstats"
)

type PixelSet struct {
	pixels []Pixel
	rgbStats [3]sstats.Summary
}

func MakePixelSet(pixels []Pixel) PixelSet {
	rgbStats := [3]sstats.Summary{ sstats.MakeSummary(), sstats.MakeSummary(), sstats.MakeSummary() }
	for _,rgb := range pixels {
		rgbStats[0].Meet( float64( rgb[0] ) )
		rgbStats[1].Meet( float64( rgb[1] ) )
		rgbStats[2].Meet( float64( rgb[2] ) )
	}
	return PixelSet{ pixels, rgbStats }
}

func (pset *PixelSet) Size() int {
	return len(pset.pixels)
}

func (pset *PixelSet) Empty() bool {
	return len(pset.pixels) == 0
}

func (pset *PixelSet) Mean() Pixel {
	return Pixel{
		uint8( math.Floor(pset.rgbStats[0].Mean() + .5) ),
		uint8( math.Floor(pset.rgbStats[1].Mean() + .5) ),
		uint8( math.Floor(pset.rgbStats[2].Mean() + .5) ),
	}
}

func (pset *PixelSet) MedianCut(axis int) (px Pixel, left PixelSet, right PixelSet) {
	l := len(pset.pixels)
	if (l < 2) {
		if l == 0 {
			return pset.Mean(), *pset, *pset
		} else { // l == 1
			return pset.pixels[0], *pset, MakePixelSet( []Pixel{} )
		}
	} else {
		px := median(pset.pixels, axis)
		if l % 2 == 0 {
			left  := MakePixelSet( pset.pixels[   :l/2] )
			right := MakePixelSet( pset.pixels[l/2:   ] )
			return px, left, right
		} else {
			left  := MakePixelSet( pset.pixels[     :l/2] )
			right := MakePixelSet( pset.pixels[l/2+1:   ] )
			return px, left, right
		}
	}
}

func (pset *PixelSet) MaxVariance() (axis int, variance float64) {
	return sstats.Max( pset.rgbStats[0].Variance(),
	                   pset.rgbStats[1].Variance(),
	                   pset.rgbStats[2].Variance() )
}

func (pset *PixelSet) MaxSpread() (axis int, spread float64) {
	return sstats.Max( pset.rgbStats[0].Max() - pset.rgbStats[0].Min(),
	                   pset.rgbStats[1].Max() - pset.rgbStats[1].Min(),
	                   pset.rgbStats[2].Max() - pset.rgbStats[2].Min() )
}

// this is adapted from github.com/giorgiga/sstats - see there for comments

func median(pixels []Pixel, axis int) Pixel {
	sz := len(pixels)
	switch sz {
		case 0:  return Pixel{}
		case 1:  return pixels[0]
		default: return findIth(pixels, axis, sz/2, sz % 2 == 0)
	}
}

func findIth(values []Pixel, axis int, i int, domean bool) Pixel {
	pi := qSortRoundR(values, axis, len(values)/2)
	if i < pi {
		return findIth(values[:pi], axis, i, domean)
	} else if pi < i {
		return findIth(values[pi:], axis, i - pi, domean)
	} else if domean {
		prev := values[0]
		for _,p := range values[:i] {
			if p[axis] > prev[axis] { prev = p }
		}
		pset := MakePixelSet( []Pixel{values[i], prev} )
		return pset.Mean()
	} else {
		return values[i]
	}
}

func qSortRoundR(values []Pixel, axis int, pivot int) int {
	values[pivot], values[0] = values[0], values[pivot]
	right := len(values) -1
	for i := right; i > 0; i-- {
		if values[i][axis] > values[0][axis] {
			values[i], values[right] = values[right], values[i]
			right--
		}
	}
	values[right], values[0] = values[0], values[right]
	return right
}
