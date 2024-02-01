package main

import "testing"

func TestHorizontal(t *testing.T) {
	flipper := NewFlipper()
	tests := map[string]string{
		"1234567890": "0987654321",
		"":           "",
		"wevaty":     "oim;yt",
		"g":          "h",
		"3 Dragons?": "8 Du;hwbl?",
		"ミaカbサc":     "ミ;カnサ,",
	}
	for test, expected := range tests {
		t.Run(test, func(t *testing.T) {
			flipped := flipper.HorizontalFlip(test)
			if flipped != expected {
				t.Fatalf("Expected %s got %s", expected, flipped)
			}
		})
	}
}

func TestVertical(t *testing.T) {
	flipper := NewFlipper()
	tests := map[string]string{
		"1234567890": "zxcvbnm,./",
		"":           "",
		"wevaty":     "sd4qgh",
		"g":          "t",
		"3 Dragons?": "c Dfqtl6w?",
		"ミaカbサc":     "ミqカ5サ3",
	}
	for test, expected := range tests {
		t.Run(test, func(t *testing.T) {
			flipped := flipper.VerticalFlip(test)
			if flipped != expected {
				t.Fatalf("Expected %s got %s", expected, flipped)
			}
		})
	}
}
