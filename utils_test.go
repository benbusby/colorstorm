package main

import (
	"testing"
)

func TestKeyChecks(t *testing.T) {
	movement := []string{
		"up",
		"down",
		"left",
		"right",
	}

	for _, m := range movement {
		isMovement := isMovementKey(m)
		isMovement = isMovement && isMovementKey("shift"+m)

		if !isMovement {
			t.Error("failed to recognize movement key")
		}
	}

	exit := []string{
		"esc",
		"q",
		"ctrl+c",
	}

	for _, e := range exit {
		isExit := isExitKey(e)

		if !isExit {
			t.Error("failed to recognize exit key")
		}
	}
}

func TestCompression(t *testing.T) {
	testData := make([]byte, 1024)
	for i := range testData {
		testData[i] = byte(i)
	}
	compressed := compress(testData)

	if len(compressed) >= len(testData) {
		t.Logf("Uncompressed length: %d\n", len(testData))
		t.Logf("Compressed length:   %d\n", len(compressed))
		t.Error("compressed data is larger than uncompressed data")
	}

	decompressed := decompress(compressed)
	if len(decompressed) != len(testData) {
		t.Error("decompression resulted in incorrect size")
	}

	for i := range decompressed {
		if decompressed[i] != testData[i] {
			t.Error("decompression resulted in incorrect data")
		}
	}
}

func TestSanitizeName(t *testing.T) {
	name := "My Theme's Cool Name"

	sn := sanitizeName(name)
	if name == sn {
		t.Error("name sanitization didn't do anything")
	} else if sn != "my_themes_cool_name" {
		t.Log("Expected: my_themes_cool_name")
		t.Logf("Actual:   %s\n", sn)
		t.Error("unexpected result from name sanitization")
	}
}

func TestHexToX256(t *testing.T) {
	// See: https://www.ditig.com/256-colors-cheat-sheet#xterm-non-system-colors
	white := hexToX256("#ffffff")   // Grey100:  231
	black := hexToX256("#000000")   // Grey0:    16
	darkRed := hexToX256("#ae0000") // ~Red3:    124 (slightly off)
	magenta := hexToX256("#ae00ae") // ~Magenta: 127 (slightly off)

	if black != 16 || white != 231 || darkRed != 124 || magenta != 127 {
		t.Logf("\nWhite:    %d (expected 231)\n"+
			"Black:    %d (expected 16)\n"+
			"Dark Red: %d (expected 124)\n"+
			"Magenta:  %d (expected 127)\n",
			white, black, darkRed, magenta)
		t.Error("Hex to X256 color conversion error, see logs")
	}
}
