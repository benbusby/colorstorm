package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/x/term"
	"github.com/lucasb-eyer/go-colorful"
	"log"
	"math"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const (
	appWidth     = 90
	appHeight    = 28
	previewWidth = 60

	colorDetailHexKey   = "color_hex"
	colorDetailFieldKey = "color_detail"

	saveDraftAction     = -1
	generateThemeAction = -2

	saveColorEditAction   = -1
	cancelColorEditAction = -2

	saveFileAction   = -1
	cancelFileAction = -2
)

var (
	theme      = newTheme()
	mainAction int
	editAction int
	saveAction int

	colorFormSelected int
	colorFormField    int

	colorEdit   = &Color{}
	colorBackup string

	hasEditedHex bool

	fileName string
)

type Color struct {
	Hex  *string
	R    uint8
	G    uint8
	B    uint8
	H    uint16
	S    uint8
	V    uint8
	init bool
}

type Styles struct {
	Base,
	HeaderText,
	Preview,
	PreviewHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

type Model struct {
	lg        *lipgloss.Renderer
	styles    *Styles
	form      *huh.Form
	width     int
	ref       Mosaic
	extraHelp []key.Binding

	colorForm *huh.Form
	saveForm  *huh.Form

	lowTermHeight  bool
	showColorTable bool
	showReference  bool
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Bold(true).
		Padding(0, 1, 0, 0)
	s.Preview = lg.NewStyle().
		PaddingLeft(1).
		MarginTop(1)
	s.PreviewHeader = lg.NewStyle().
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

func NewModel(refMosaic Mosaic, lowTermHeight bool) Model {
	m := Model{
		width:         appWidth,
		lowTermHeight: lowTermHeight,
		showReference: !lowTermHeight,
	}

	if len(refMosaic.Image) > 0 {
		m.ref = refMosaic
	}

	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	m.form = createForm(m.lg)

	if len(refMosaic.Image) > 0 {
		m.extraHelp = []key.Binding{
			key.NewBinding(
				key.WithKeys("`"),
				key.WithHelp("`", "show/hide ref")),
		}

		if lowTermHeight {
			m.extraHelp = append(m.extraHelp, key.NewBinding(
				key.WithKeys("~"),
				key.WithHelp("~", "toggle ref mode")))
		}
	}

	return m
}

func createForm(lg *lipgloss.Renderer) *huh.Form {
	themeName := huh.NewInput().Title("Theme Name").Placeholder("My Theme").Value(&theme.Name)

	// TODO: Determine light vs dark theme colors and adjust foreground color
	bgLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#fff")).
		Background(lipgloss.Color(*theme.Background)).
		Render(fmt.Sprintf("Background [%s]", *theme.Background))

	fgLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color(*theme.Foreground)).
		Render(fmt.Sprintf("Foreground [%s]", *theme.Foreground))

	funcLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color(*theme.Function)).
		Render(fmt.Sprintf("Function   [%s]", *theme.Function))

	constLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color(*theme.Constant)).
		Render(fmt.Sprintf("Constant   [%s]", *theme.Constant))

	keywordLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color(*theme.Keyword)).
		Render(fmt.Sprintf("Keyword    [%s]", *theme.Keyword))

	commentLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#fff")).
		Background(lipgloss.Color(*theme.Comment)).
		Render(fmt.Sprintf("Comment    [%s]", *theme.Comment))

	numberLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color(*theme.Number)).
		Render(fmt.Sprintf("Number     [%s]", *theme.Number))

	stringLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color(*theme.String)).
		Render(fmt.Sprintf("String     [%s]", *theme.String))

	typeLabel := lg.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color(*theme.Type)).
		Render(fmt.Sprintf("Type       [%s]", *theme.Type))

	themeSelect := huh.NewSelect[int]().
		Options(
			huh.NewOption(bgLabel, BackgroundIndex),
			huh.NewOption(fgLabel, ForegroundIndex),
			huh.NewOption(funcLabel, FunctionIndex),
			huh.NewOption(constLabel, ConstantIndex),
			huh.NewOption(keywordLabel, KeywordIndex),
			huh.NewOption(commentLabel, CommentIndex),
			huh.NewOption(numberLabel, NumberIndex),
			huh.NewOption(stringLabel, StringIndex),
			huh.NewOption(typeLabel, TypeIndex),
			huh.NewOption("Save Draft", saveDraftAction),
			huh.NewOption("Generate Theme", generateThemeAction),
		).Value(&mainAction)

	form := huh.NewForm(huh.NewGroup(themeName, themeSelect)).
		WithWidth(25).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(huh.ThemeCatppuccin())

	if mainAction > 0 {
		themeSelect.Focus()
	} else {
		themeName.Focus()
	}

	return form
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func createSaveForm() *huh.Form {
	return huh.NewForm(huh.NewGroup(
		huh.NewInput().Title("Save As...").Placeholder("my_theme.json").Value(&fileName),
		huh.NewSelect[int]().
			Options(
				huh.NewOption("Save File", saveFileAction),
				huh.NewOption("Cancel", cancelFileAction)).
			Value(&saveAction))).
		WithWidth(25).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(huh.ThemeCatppuccin())
}

func createColorForm(colorName string, hexColor *string) *huh.Form {
	color, err := colorful.Hex(*hexColor)
	if err == nil {
		h, s, v := color.Hsv()
		if !colorEdit.init || hasEditedHex {
			colorEdit.S = uint8(roundFloat(s, 2) * 100)
			colorEdit.V = uint8(roundFloat(v, 2) * 100)
			hasEditedHex = false
		}

		colorEdit.Hex = hexColor
		colorEdit.H = uint16(roundFloat(h, 2))
		colorEdit.R = uint8(color.R * 255)
		colorEdit.G = uint8(color.G * 255)
		colorEdit.B = uint8(color.B * 255)
		colorEdit.init = true
	} else {
		_ = os.WriteFile("err.out", []byte("ERROR"), 777)
	}

	rLabel := fmt.Sprintf("R: [%03d / 255]", colorEdit.R)
	gLabel := fmt.Sprintf("G: [%03d / 255]", colorEdit.G)
	bLabel := fmt.Sprintf("B: [%03d / 255]", colorEdit.B)
	hLabel := fmt.Sprintf("H: [%03d / 360]", colorEdit.H)
	sLabel := fmt.Sprintf("S: [%03d / 100]", colorEdit.S)
	vLabel := fmt.Sprintf("V: [%03d / 100]", colorEdit.V)

	input := huh.NewInput().Title(colorName).Value(hexColor).Key(colorDetailHexKey)
	colorDetails := huh.NewSelect[int]().
		Title("RGB/HSV").
		Description("→ = +1\n← = -1\nshift + → = +10\nshift + ← = -10").
		Options(
			huh.NewOption(rLabel, 0),
			huh.NewOption(gLabel, 1),
			huh.NewOption(bLabel, 2),
			huh.NewOption(hLabel, 3),
			huh.NewOption(sLabel, 4),
			huh.NewOption(vLabel, 5)).
		Value(&colorFormField).
		Key(colorDetailFieldKey)

	form := huh.NewForm(
		huh.NewGroup(
			input,
			colorDetails,
			huh.NewSelect[int]().
				Options(
					huh.NewOption("Save Changes", saveColorEditAction),
					huh.NewOption("Cancel", cancelColorEditAction)).
				Value(&editAction),
		)).
		WithWidth(25).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(huh.ThemeCatppuccin())

	if colorFormSelected == 0 {
		for form.GetFocusedField().GetKey() != colorDetailHexKey {
			form.NextField()
		}
	} else if colorFormSelected == 1 {
		for form.GetFocusedField().GetKey() != colorDetailFieldKey {
			form.NextField()
		}
	}

	return form
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, appWidth) -
			m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "~":
			m.showColorTable = !m.showColorTable
			return m, nil
		case "`":
			m.showReference = !m.showReference
			return m, nil
		case "esc", "q", "ctrl+c":
			return m, tea.Quit
		case "shift+left", "left", "shift+right", "right":
			if m.colorForm != nil && m.colorForm.GetFocusedField().GetKey() == colorDetailFieldKey {
				val := 1

				if strings.HasPrefix(msg.String(), "shift") {
					val *= 10
				}

				if strings.HasSuffix(msg.String(), "left") {
					val *= -1
				}

				switch colorFormField {
				case 0:
					colorEdit.R += uint8(val)
				case 1:
					colorEdit.G += uint8(val)
				case 2:
					colorEdit.B += uint8(val)
				case 3:
					colorEdit.H += uint16(val)
				case 4:
					colorEdit.S += uint8(val)
				case 5:
					colorEdit.V += uint8(val)
				}

				var c colorful.Color
				if colorFormField >= 3 {
					c = colorful.Hsv(
						float64(colorEdit.H),
						roundFloat(float64(colorEdit.S)/100.0, 2),
						roundFloat(float64(colorEdit.V)/100.0, 2))
				} else {
					c = colorful.Color{
						R: float64(colorEdit.R) / 255.0,
						G: float64(colorEdit.G) / 255.0,
						B: float64(colorEdit.B) / 255.0,
					}
				}

				if colorFormField < 4 {
					hasEditedHex = true
				}

				newBG := c.Hex()
				theme.SetHexColor(mainAction, &newBG)
				colorFormSelected = 1
				m.colorForm = createColorForm(theme.GetColorName(mainAction), theme.GetHexColor(mainAction))
			}
		case "tab", "enter":
			if m.colorForm != nil {
				if len(*theme.Background) == 7 && m.colorForm.GetFocusedField().GetKey() == colorDetailHexKey {
					colorFormSelected = 1
					hasEditedHex = true
					m.colorForm = createColorForm(theme.GetColorName(mainAction), theme.GetHexColor(mainAction))
					return m, nil
				}
			}
		}
	}

	var cmds []tea.Cmd

	// Process the form
	var (
		form tea.Model
		cmd  tea.Cmd
	)

	if m.colorForm != nil {
		form, cmd = m.colorForm.Update(msg)
	} else if m.saveForm != nil {
		form, cmd = m.saveForm.Update(msg)
	} else {
		form, cmd = m.form.Update(msg)
	}

	if f, ok := form.(*huh.Form); ok {
		if m.colorForm != nil {
			m.colorForm = f
		} else if m.saveForm != nil {
			m.saveForm = f
		} else {
			m.form = f
		}
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		if mainAction == saveDraftAction {
			m.form = createForm(m.lg)
			m.saveForm = createSaveForm()
			cmds = append(cmds, m.saveForm.Init(), tea.WindowSize())
		} else if mainAction == generateThemeAction {
			cmds = append(cmds, tea.Quit)
		} else if m.colorForm == nil {
			m.form = createForm(m.lg)
			m.colorForm = createColorForm(
				theme.GetColorName(mainAction),
				theme.GetHexColor(mainAction))
			colorBackup = *theme.GetHexColor(mainAction)
			cmds = append(cmds, m.colorForm.Init(), tea.WindowSize())
		}
	} else if m.colorForm != nil && m.colorForm.State == huh.StateCompleted {
		if editAction == cancelColorEditAction {
			theme.SetHexColor(mainAction, &colorBackup)
		}

		colorFormSelected = 0
		m.form = createForm(m.lg)
		m.colorForm = nil
		cmds = append(cmds, m.form.Init(), tea.WindowSize(), m.form.NextField())
	} else if m.saveForm != nil && m.saveForm.State == huh.StateCompleted {
		if saveAction == saveFileAction {
			refBytes, _ := SerializeMosaic(m.ref)
			theme.Reference = refBytes

			themeBytes, err := json.Marshal(theme)
			if err != nil {
				log.Println("Error serializing theme", err)
			}

			err = os.WriteFile(fileName, themeBytes, 0666)
			if err != nil {
				log.Println("Error writing file", err)
			}

			cmds = append(cmds, tea.Quit)
		} else {
			m.form = createForm(m.lg)
			m.saveForm = nil
			cmds = append(cmds, m.form.Init(), tea.WindowSize(), m.form.NextField())
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	header := m.appBoundaryView("COLORSTORM")
	keyBinds := append(m.form.KeyBinds(), m.extraHelp...)
	footer := m.appBoundaryView(m.form.Help().ShortHelpView(keyBinds))

	var colorTable string
	if m.showColorTable {
		colorTable = formatColorTable(m.ref.Colors)
	}

	if m.lowTermHeight && m.showReference {
		if m.showColorTable {
			return s.Base.Render(
				header + "\n" +
					colorTable + "\n" + footer)
		} else {
			return s.Base.Render(
				header + "\n" +
					m.ref.Image + "\n" + footer)
		}
	}

	// Form
	var formView string
	if m.colorForm != nil {
		formView = strings.TrimSuffix(m.colorForm.View(), "\n")
	} else if m.saveForm != nil {
		formView = strings.TrimSuffix(m.saveForm.View(), "\n")
	} else {
		formView = strings.TrimSuffix(m.form.View(), "\n")
	}

	form := m.lg.NewStyle().Margin(1, 0).Render(formView)

	statusMarginLeft := m.width -
		previewWidth -
		lipgloss.Width(form) -
		s.Preview.GetMarginRight()

	// Status
	var preview string
	{
		preview = s.Preview.
			Height(lipgloss.Height(form)).
			Width(previewWidth).
			MarginLeft(statusMarginLeft).
			Render(s.PreviewHeader.Render("Preview") + "\n" +
				parseSyntaxHighlighting(sampleText, theme))
	}

	// Reference
	var reference string
	if len(m.ref.Image) > 0 && m.showReference {
		if m.showColorTable {
			reference = colorTable
		} else {
			reference = m.ref.Image
		}
	}

	body := lipgloss.JoinHorizontal(lipgloss.Top, form, preview)
	if len(reference) > 0 {
		body = lipgloss.JoinVertical(lipgloss.Top, reference, body)
	}

	return s.Base.Render(header + "\n" + body + "\n" + footer)
}

func (m Model) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
	)
}

func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
	)
}

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
	}

	_, height, _ := term.GetSize(0)
	lowTermHeight := height < appHeight*2

	m := NewModel(refMosaic, lowTermHeight)
	_, err = tea.NewProgram(m, tea.WithAltScreen()).Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if saveAction == saveFileAction {
		fmt.Printf("Theme file saved to: %s\n", fileName)
	}
}
