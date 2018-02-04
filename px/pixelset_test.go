package px

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	pset := MakePixelSet( []Pixel{} )
	if !pset.Empty() { t.Errorf("expected empty") }
	if pset.Size() != 0 { t.Errorf("expected size 0") }
	// the only thing we care is that these does not explode
	pset.Mean()
	pset.MaxSpread()
	pset.MaxVariance()
	pset.MedianCut(0)
	pset.MedianCut(1)
	pset.MedianCut(2)
}

func TestSingle(t *testing.T) {
	pixels := MakePixelSet( []Pixel{  Pixel{ 0, 0, 0 } } )

	mean := fmt.Sprintf("%v", pixels.Mean())
	if "000000" != mean { t.Errorf("expected black mean but got %v", mean) }

	mpx, l, r := pixels.MedianCut(0)

	smpx := fmt.Sprintf("%v", mpx)
	if "000000" != smpx { t.Errorf("expected black median but got %v", smpx) }
	if l.Size() != 1 { t.Errorf("expected left of size 1 but got %v", l.Size()) }
	if r.Size() != 0 { t.Errorf("expected right of size 0 but got %v", r.Size()) }
}

func TestTwo(t *testing.T) {
	pixels := MakePixelSet( []Pixel{  Pixel{ 0, 0, 0 }, Pixel{ 255, 255, 255 } } )

	mean := fmt.Sprintf("%v", pixels.Mean())
	if "808080" != mean { t.Errorf("expected 7f7f7f mean but got %v", mean) }

	mpx, l, r := pixels.MedianCut(1)

	smpx := fmt.Sprintf("%v", mpx)
	if "808080" != smpx { t.Errorf("expected 7f7f7f median but got %v", smpx) }
	if l.Size() != 1 { t.Errorf("expected left of size 1 but got %v", l.Size()) }
	if r.Size() != 1 { t.Errorf("expected right of size 1 but got %v", r.Size()) }
}

func TestThree(t *testing.T) {
	pixels := MakePixelSet( []Pixel{  Pixel{ 0, 0, 0 }, Pixel{ 64, 64, 64 }, Pixel{ 255, 255, 255 } } )

	mean := fmt.Sprintf("%v", pixels.Mean())
	if "6a6a6a" != mean { t.Errorf("expected 6a6a6a mean but got %v", mean) }

	mpx, l, r := pixels.MedianCut(1)

	smpx := fmt.Sprintf("%v", mpx)
	if "404040" != smpx { t.Errorf("expected 404040 median but got %v", smpx) }
	if l.Size() != 1 { t.Errorf("expected left of size 1 but got %v", l.Size()) }
	if r.Size() != 1 { t.Errorf("expected right of size 1 but got %v", r.Size()) }
}
