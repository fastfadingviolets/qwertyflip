package flip_test

import (
	"testing"

	"github.com/fastfadingviolets/qwertyflip/flip"
)

func TestHorizontal(t *testing.T) {
	flipper := flip.NewFlipper()
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
	flipper := flip.NewFlipper()
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

func TestShift(t *testing.T) {
	flipper := flip.NewFlipper()
	type test struct {
		line string
		by   int
	}
	tests := map[test]string{
		{"1234567890", 3}:  "4567890qwe",
		{"", 500}:          "",
		{"3 Dragons?", -8}: "b D6euqkr?",
		{"1", -1}:          "/",
		{"/", 1}:           "1",
		{"food", 40}:       "food",
		{"food", 44}:       "kddj",
		{"food", 97}:       "1nn/",
		{"ミaカbサc", -28}:    "ミcカ7サ5",
	}
	for test, expected := range tests {
		t.Run(test.line, func(t *testing.T) {
			shifted := flipper.Shift(test.by, test.line)
			if shifted != expected {
				t.Fatalf("Expected %s got %s", expected, shifted)
			}
		})
	}
}

func TestCommands(t *testing.T) {
	flipper := flip.NewFlipper()
	type test struct {
		line    string
		command string
	}
	tests := map[test]string{
		{"1234567890", "H,V"}:    "/.,mnbvcxz",
		{"Hello World", "2"}:     "Htzza Wayzg",
		{"Qwerty", "5,H"}:        "Qrewq;",
		{"asdf", "V,-3"}:         "890q",
		{"3 Dragons?", "V,-8,H"}: "h Dt84;c7?",
		{"ミaカbサc", "-28,H"}:      "ミ,カ4サ6",
	}
	for test, expected := range tests {
		t.Run(test.line+"_"+test.command, func(t *testing.T) {
			result, err := flipper.RunCommand(test.command, test.line)
			if err != nil {
				t.Fatal(err)
			}
			if result != expected {
				t.Fatalf("Expected %s got %s", expected, result)
			}
		})
	}
}
