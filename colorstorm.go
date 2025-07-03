package main

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const (
	appWidth     = 90
	appHeight    = 28
	previewWidth = 65
)

var (
	theme      = newDefaultTheme()
	mainAction int

	colorFormSelected int
	colorFormField    int

	colorEdit   = &Color{}
	colorBackup string

	updateSV bool

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
	PreviewHeader lipgloss.Style
}

type Model struct {
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
	ref    Mosaic
	error  string

	refVisibilityHelp key.Binding
	refModeHelp       key.Binding
	pickerHelp        key.Binding

	colorForm       *huh.Form
	saveColorAction *bool

	saveForm       *huh.Form
	saveFileAction *bool

	generatorForm   *huh.Form
	generatorValues *GeneratorFormValues

	lowTermHeight  bool
	showColorTable bool
	showReference  bool

	showColorPicker  bool
	highlightedColor string
	backupColor      string
	colorPickerX     int
	colorPickerY     int

	output       []string
	outputErrors []string
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

	m.refVisibilityHelp = key.NewBinding(
		key.WithKeys("`"),
		key.WithHelp("`", "show/hide ref"))

	m.refModeHelp = key.NewBinding(
		key.WithKeys("~"),
		key.WithHelp("~", "toggle ref mode"))

	m.pickerHelp = key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "color picker"))

	return m
}

func (m *Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, appWidth) -
			m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		if m.showColorPicker &&
			!isMovementKey(msg.String()) &&
			!isExitKey(msg.String()) &&
			msg.String() != "enter" {
			return m, nil
		}

		if msg.String() != "~" &&
			msg.String() != "`" &&
			!isExitKey(msg.String()) &&
			m.showReference &&
			m.lowTermHeight {
			// Ignore keyboard input when viewing reference in a
			// small terminal window
			return m, nil
		}

		switch msg.String() {
		case "p":
			if len(m.ref.Image) > 0 &&
				mainAction >= BackgroundIndex &&
				m.form.GetFocusedField().GetKey() == themeActionKey {
				m.showColorPicker = true
				m.backupColor = *theme.GetHexColor(mainAction)
				_, m.highlightedColor = m.ref.HighlightPixel(
					m.colorPickerX,
					m.colorPickerY)
				theme.SetHexColor(mainAction, m.highlightedColor)
				m.form = createForm(m.lg)
				return m, tea.Batch(m.form.Init(), m.form.NextField())
			}
		case "~":
			m.showColorTable = !m.showColorTable
			return m, nil
		case "`":
			m.showReference = !m.showReference
			return m, nil
		case "esc", "q", "ctrl+c":
			if m.showColorPicker {
				m.showColorPicker = false
				theme.SetHexColor(mainAction, m.backupColor)
				m.form = createForm(m.lg)
				return m, tea.Batch(m.form.Init(), m.form.NextField())
			}
			return m, tea.Quit
		case "shift+left", "left", "shift+right", "right", "shift+up", "up", "shift+down", "down":
			if m.showColorPicker {
				// Update color picker coordinates
				m.colorPickerX, m.colorPickerY = updatePickerCoords(
					msg.String(),
					m.colorPickerX,
					m.colorPickerY)
				_, m.highlightedColor = m.ref.HighlightPixel(m.colorPickerX, m.colorPickerY)
				theme.SetHexColor(mainAction, m.highlightedColor)
				m.form = createForm(m.lg)
				return m, tea.Batch(m.form.Init(), m.form.NextField())
			} else if m.colorForm != nil && m.colorForm.GetFocusedField().GetKey() == colorDetailFieldKey {
				// Handle arrow key modification of color fields
				newColor := updateFocusedColorField(msg.String())
				theme.SetHexColor(mainAction, newColor)
				colorFormSelected = 1
				m.colorForm, m.saveColorAction = createColorForm(
					theme.GetColorName(mainAction),
					theme.GetHexColor(mainAction))
			}
		case "tab", "enter":
			if m.showColorPicker && msg.String() == "enter" {
				// Select the currently highlighted color from
				// the color picker
				theme.SetHexColor(mainAction, m.highlightedColor)
				m.showColorPicker = false
				return m, nil
			} else if m.colorForm != nil {
				// Update RGB/HSV values
				if len(*theme.Background) == 7 && m.colorForm.GetFocusedField().GetKey() == colorDetailHexKey {
					colorFormSelected = 1
					updateSV = true
					m.colorForm, m.saveColorAction = createColorForm(
						theme.GetColorName(mainAction),
						theme.GetHexColor(mainAction))
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
	} else if m.generatorForm != nil {
		form, cmd = m.generatorForm.Update(msg)
	} else {
		form, cmd = m.form.Update(msg)
	}

	if f, ok := form.(*huh.Form); ok {
		if m.colorForm != nil {
			m.colorForm = f
		} else if m.saveForm != nil {
			m.saveForm = f
		} else if m.generatorForm != nil {
			m.generatorForm = f
		} else {
			m.form = f
		}
		cmds = append(cmds, cmd)
	}

	formCmds := m.updateForms()
	cmds = append(cmds, formCmds...)

	return m, tea.Batch(cmds...)
}

func updatePickerCoords(msg string, x, y int) (int, int) {
	xMod := 0
	yMod := 0

	if strings.HasSuffix(msg, "left") {
		yMod = -1
	} else if strings.HasSuffix(msg, "right") {
		yMod = 1
	} else if strings.HasSuffix(msg, "up") {
		xMod = -1
	} else if strings.HasSuffix(msg, "down") {
		xMod = 1
	}

	if strings.HasPrefix(msg, "shift") {
		xMod *= 10
		yMod *= 10
	}

	return max(x+xMod, 0), max(y+yMod, 0)
}

func updateFocusedColorField(msg string) string {
	val := 0

	if strings.HasSuffix(msg, "left") || strings.HasSuffix(msg, "right") {
		val = 1
	}

	if strings.HasPrefix(msg, "shift") {
		val *= 10
	}

	if strings.HasSuffix(msg, "left") {
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
		// Saturation and value need to be updated for R/G/B/H changes
		updateSV = true
	}

	return c.Hex()
}

func (m *Model) updateForms() []tea.Cmd {
	if m.form.State == huh.StateCompleted {
		m.error = ""
		if mainAction == saveDraftAction {
			m.form = createForm(m.lg)
			m.saveForm, m.saveFileAction = createSaveForm()
			return []tea.Cmd{m.saveForm.Init(), tea.WindowSize()}
		} else if mainAction == generateThemeAction {
			m.form = createForm(m.lg)
			err := theme.Validate()
			if err != nil {
				m.error = err.Error()
				return []tea.Cmd{m.form.Init(), tea.WindowSize()}
			}

			m.generatorForm, m.generatorValues = createGeneratorForm()
			return []tea.Cmd{m.generatorForm.Init(), tea.WindowSize()}
		} else if m.colorForm == nil {
			m.form = createForm(m.lg)
			m.colorForm, m.saveColorAction = createColorForm(
				theme.GetColorName(mainAction),
				theme.GetHexColor(mainAction))
			colorBackup = *theme.GetHexColor(mainAction)
			return []tea.Cmd{m.colorForm.Init(), tea.WindowSize()}
		}
	} else if m.colorForm != nil && m.colorForm.State == huh.StateCompleted {
		return m.updateColorForm()
	} else if m.saveForm != nil && m.saveForm.State == huh.StateCompleted {
		return m.updateSaveForm()
	} else if m.generatorForm != nil && m.generatorForm.State == huh.StateCompleted {
		return m.updateGeneratorForm()
	}

	return []tea.Cmd{}
}

func (m *Model) updateGeneratorForm() []tea.Cmd {
	if m.generatorValues.Confirmed {
		final := theme.Finalize(*m.generatorValues)
		for _, k := range m.generatorValues.Editors {
			generatorFn, ok := generatorMap[k]
			if !ok {
				continue
			}

			themeFile, err := generatorFn(final)
			if err != nil {
				m.outputErrors = append(
					m.outputErrors,
					err.Error())
				continue
			}

			err = os.WriteFile(themeFile.FileName, themeFile.Contents, 0644)
			if err != nil {
				m.outputErrors = append(m.outputErrors, err.Error())
			} else {
				msg := fmt.Sprintf("%s: %s", k, themeFile.FileName)
				m.output = append(m.output, msg)
			}
		}
		return []tea.Cmd{tea.Quit}
	} else {
		m.form = createForm(m.lg)
		m.generatorForm = nil
		return []tea.Cmd{m.form.Init(), tea.WindowSize()}
	}
}

func (m *Model) updateColorForm() []tea.Cmd {
	if !*m.saveColorAction {
		// Revert back to original hex color
		theme.SetHexColor(mainAction, colorBackup)
	}

	colorFormSelected = 0
	m.form = createForm(m.lg)
	m.colorForm = nil
	return []tea.Cmd{m.form.Init(), tea.WindowSize(), m.form.NextField()}
}

func (m *Model) updateSaveForm() []tea.Cmd {
	if *m.saveFileAction {
		refBytes := SerializeMosaic(m.ref)
		theme.Reference = refBytes

		themeBytes, err := json.Marshal(theme)
		if err != nil {
			log.Println("Error serializing theme", err)
		}

		err = os.WriteFile(fileName, themeBytes, 0644)
		if err != nil {
			log.Println("Error writing file", err)
		}

		return []tea.Cmd{tea.Quit}
	} else {
		m.form = createForm(m.lg)
		m.saveForm = nil
		return []tea.Cmd{m.form.Init(), tea.WindowSize(), m.form.NextField()}
	}
}

func (m *Model) View() string {
	s := m.styles

	header := m.appBoundaryView("COLORSTORM")
	footer := m.appBoundaryView(m.helpView())

	var colorTable string
	if m.showColorTable {
		colorTable = formatColorTable(m.ref.Colors)
	}

	if m.lowTermHeight && (m.showReference || m.showColorPicker) {
		var body string
		if m.showColorPicker {
			hlRef, hex := m.ref.HighlightPixel(m.colorPickerX, m.colorPickerY)
			fg := lipgloss.Color(gamut.ToHex(gamut.Contrast(gamut.Hex(hex))))
			bg := lipgloss.Color(hex)
			body = hlRef + "\n" + m.lg.NewStyle().
				Foreground(fg).
				Background(bg).
				Render(fmt.Sprintf("Color: %s", hex))
		} else {
			body = m.ref.Image
			if m.showColorTable {
				body += "\n" + colorTable
			}
		}

		return s.Base.Render(header + "\n" + body + "\n" + footer)
	}

	// Form
	var formView string
	if m.colorForm != nil {
		formView = strings.TrimSuffix(m.colorForm.View(), "\n")
	} else if m.saveForm != nil {
		formView = strings.TrimSuffix(m.saveForm.View(), "\n")
	} else if m.generatorForm != nil {
		formView = strings.TrimSuffix(m.generatorForm.View(), "\n")
	} else {
		formView = strings.TrimSuffix(m.form.View(), "\n")
	}

	form := m.lg.NewStyle().Margin(1, 0).Render(formView)
	statusMarginLeft := m.width -
		previewWidth -
		lipgloss.Width(form) -
		s.Preview.GetMarginRight()

	// Status
	preview := s.Preview.
		Height(lipgloss.Height(form)).
		Width(previewWidth).
		MarginLeft(statusMarginLeft).
		Render(s.PreviewHeader.Render("Preview") + "\n" +
			parseSyntaxHighlighting(sampleText, theme))

	// Reference
	var reference string
	if len(m.ref.Image) > 0 && m.showReference {
		if m.showColorPicker {
			hlRef, hex := m.ref.HighlightPixel(m.colorPickerX, m.colorPickerY)
			fg := lipgloss.Color(gamut.ToHex(gamut.Contrast(gamut.Hex(hex))))
			bg := lipgloss.Color(hex)
			reference = hlRef + "\n" + m.lg.NewStyle().
				Foreground(fg).
				Background(bg).
				Render(fmt.Sprintf("Color: %s", hex))
		} else {
			reference = m.ref.Image
			if m.showColorTable {
				reference += "\n" + colorTable
			}
		}
	}

	body := lipgloss.JoinHorizontal(lipgloss.Top, form, preview)
	if len(reference) > 0 {
		body = lipgloss.JoinVertical(lipgloss.Top, reference, body)
	}

	return s.Base.Render(header + "\n" + body + "\n" + footer)
}

func (m *Model) helpView() string {
	if len(m.error) > 0 {
		return huh.ThemeCatppuccin().Focused.ErrorMessage.Render(m.error)
	}

	var (
		footer       string
		helpElements []key.Binding
	)

	if len(m.ref.Image) > 0 {
		helpElements = []key.Binding{m.refVisibilityHelp}
		if m.showReference {
			helpElements = append(helpElements, m.refModeHelp)
		}

		if mainAction >= BackgroundIndex && m.form.GetFocusedField().GetKey() == themeActionKey {
			helpElements = append(helpElements, m.pickerHelp)
		}
	}

	if m.generatorForm != nil && m.generatorForm.GetFocusedField().GetKey() == editorSelectKey {
		helpElements = append(helpElements, key.NewBinding(
			key.WithKeys("space"),
			key.WithHelp("space", "select / deselect")))
	}

	for _, binding := range helpElements {
		helpKey := fmt.Sprintf("[%s] %s ",
			huh.ThemeCatppuccin().Help.ShortKey.Render(
				strings.Join(binding.Keys(), ",")),
			huh.ThemeCatppuccin().Help.ShortDesc.Render(
				binding.Help().Desc))
		footer += helpKey
	}

	return footer
}

func (m *Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
	)
}
