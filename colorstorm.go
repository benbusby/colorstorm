package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/x/term"
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
	previewWidth = 60
)

type action int

const (
	saveDraft action = iota
	generateTheme
)

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
	action    action
	ref       ReferenceMosaic
	extraHelp []key.Binding

	lowTermHeight     bool
	showColorTable    bool
	showFullReference bool
}

var theme = newTheme()

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

func NewModel(refMosaic ReferenceMosaic, lowTermHeight bool) Model {
	m := Model{
		width:             appWidth,
		lowTermHeight:     lowTermHeight,
		showFullReference: !lowTermHeight,
	}

	if len(refMosaic.Image) > 0 {
		m.ref = refMosaic
	}

	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Theme Name").Placeholder("My Theme"),
			huh.NewInput().Title("Background").Value(&theme.Background),
			huh.NewInput().Title("Foreground").Value(&theme.Foreground),
			huh.NewInput().Title("Function").Value(&theme.Function),
			huh.NewInput().Title("Constant").Value(&theme.Constant),
			huh.NewInput().Title("Keyword").Value(&theme.Keyword),
			huh.NewInput().Title("Comment").Value(&theme.Comment),
			huh.NewInput().Title("Number").Value(&theme.Number),
			huh.NewInput().Title("String").Value(&theme.String),
			huh.NewInput().Title("Type").Value(&theme.Type),

			huh.NewSelect[action]().
				Options(
					huh.NewOption("Save Draft", saveDraft),
					huh.NewOption("Generate Theme", generateTheme),
				).
				Value(&m.action),
		),
	).WithWidth(25).WithShowHelp(false).WithShowErrors(false).WithTheme(huh.ThemeCatppuccin())

	if len(refMosaic.Image) > 0 {
		m.extraHelp = []key.Binding{
			key.NewBinding(
				key.WithKeys("`"),
				key.WithHelp("`", "show/hide ref")),
		}

		if !lowTermHeight {
			m.extraHelp = append(m.extraHelp, key.NewBinding(
				key.WithKeys("~"),
				key.WithHelp("~", "toggle ref mode")))
		}
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, appWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "~":
			m.showColorTable = !m.showColorTable
			return m, nil
		case "`":
			m.showFullReference = !m.showFullReference
			return m, nil
		case "esc", "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	header := m.appBoundaryView("COLORSTORM")
	keyBinds := append(m.form.KeyBinds(), m.extraHelp...)
	footer := m.appBoundaryView(m.form.Help().ShortHelpView(keyBinds))

	if m.lowTermHeight && m.showFullReference {
		return s.Base.Render(
			header + "\n" +
				m.ref.Image + "\n" +
				m.ref.ColorTable + "\n" + footer)
	}

	// Form
	formView := strings.TrimSuffix(m.form.View(), "\n")
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
	if len(m.ref.Image) > 0 && !m.lowTermHeight && m.showFullReference {
		if m.showColorTable {
			reference = m.ref.ColorTable
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
		refMosaic ReferenceMosaic
	)

	if len(os.Args) > 1 {
		refMosaic, err = generateReferenceMosaic(os.Args[1], appWidth, appHeight)
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
		os.Exit(1)
	}
}
