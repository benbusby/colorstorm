package main

import (
	"encoding/json"
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		err       error
		refMosaic Mosaic
	)

	draftImport := flag.String("i", "", "import saved theme json")
	refImg := flag.String("r", "", "jpg or png reference image")
	quantize := flag.Int("q", 50, "color quantization amount [0-255]")
	flag.Parse()

	if len(*draftImport) > 0 {
		f, err := os.ReadFile(*draftImport)
		if err != nil {
			log.Fatalln("Error importing theme", err)
		}

		err = json.Unmarshal(f, theme)
		if err != nil {
			log.Fatalln("Error reading theme file", err)
		}

		if len(theme.Reference) > 0 {
			refMosaic, err = DeserializeMosaic(theme.Reference)
		}
	}

	if len(*refImg) > 0 {
		refMosaic, err = GenerateQuantizedMosaic(*refImg, appWidth, appHeight, *quantize)
		if err != nil {
			log.Fatalln("Error generating mosaic", err)
			return
		}
	} else {
		theme = newRandomTheme()
	}

	_, height, _ := term.GetSize(0)
	lowTermHeight := height < appHeight*2

	m := NewModel(refMosaic, lowTermHeight)
	_, err = tea.NewProgram(&m, tea.WithAltScreen()).Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if m.saveFileAction != nil && *m.saveFileAction {
		fmt.Printf("Theme file saved to: %s\n", fileName)
	} else if len(m.output) > 0 {
		msg := fmt.Sprintf("Output:\n- %s", strings.Join(m.output, "\n- "))
		fmt.Println(msg)
	}

	if len(m.outputErrors) > 0 {
		msg := fmt.Sprintf("Error:%s", strings.Join(m.outputErrors, "\n! "))
		fmt.Println(msg)
	}
}
