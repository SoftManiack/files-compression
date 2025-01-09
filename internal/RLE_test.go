package internal

import (
	"testing"
)

func TestRLECompression(t *testing.T) {

	// input [ 1,1,1,255,255,255,255,10,90,70,70,70,70 ]
	// output [ 1,3,255,4,10,1,90,1,70,4 ]

	input := []byte{1, 1, 1, 255, 255, 255, 255, 10, 90, 70, 70, 70, 70}

	output := internal.LRECompression(input)

	if output != []byte{1, 3, 255, 4, 10, 1, 90, 1, 70, 4} {
		t.Errorf("compression error")
	}
}
