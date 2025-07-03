package main

import (
	"encoding/json"
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"github.com/muesli/gamut"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		err       error
		refMosaic Mosaic
	)

	themeFile := flag.String("f", "", "load theme json file")

	refImg := flag.String("i", "", "jpg or png reference image")
	quantize := flag.Int("q", 0, "color quantization amount [0-255]")
	lightTheme := flag.Bool("l", false, "create a light theme")
	initNoColor := flag.Bool("x", false, "initialize without any colors")
	monochrome := flag.Bool("m", false, "generate monochrome theme")
	seedColor := flag.String("c", "", "seed color (hex)")
	flag.Parse()

	if len(*themeFile) > 0 {
		f, err := os.ReadFile(*themeFile)
		if err != nil {
			log.Fatalln("Error importing theme", err)
		}

		err = json.Unmarshal(f, theme)
		if err != nil {
			log.Fatalln("Error reading theme file", err)
		}

		if len(theme.Reference) > 0 {
			refMosaic = DeserializeMosaic(theme.Reference)
		}
	} else if len(*refImg) > 0 {
		refMosaic, err = GenerateQuantizedMosaic(*refImg, appWidth-5, appHeight, *quantize)
		if err != nil {
			log.Fatalln("Error generating mosaic", err)
			return
		}

		theme = newNoColorTheme(*lightTheme)
	} else {
		if *initNoColor {
			theme = newNoColorTheme(*lightTheme)
		} else if *monochrome && len(*seedColor) > 0 {
			theme = newMonoTheme(*lightTheme, gamut.Hex(*seedColor))
		} else if *monochrome {
			theme = newRandomMonoTheme(*lightTheme)
		} else {
			theme = newRandomTheme(*lightTheme)
		}
	}

	_, height, _ := term.GetSize(0)
	lowTermHeight := height < appHeight+(refMosaic.Height/2)

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
