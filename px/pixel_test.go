package px

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {

	px := Pixel { 0, 10, 255 }

	if "000aff" != fmt.Sprintf("%x", px) { t.Errorf("expected %v but got %v", "000aff", fmt.Sprintf("%x", px)) }
	if "000AFF" != fmt.Sprintf("%X", px) { t.Errorf("expected %v but got %v", "000AFF", fmt.Sprintf("%X", px)) }

	if "000aff" != fmt.Sprintf("%v", px) { t.Errorf("expected %v but got %v", "000aff", fmt.Sprintf("%v", px))  }

	if "px.Pixel { 0, 10, 255 }" != fmt.Sprintf("%#v", px) { t.Errorf("expected %v but got %v", "px.px { 0, 10, 255 }", fmt.Sprintf("%#v", px)) }

}
