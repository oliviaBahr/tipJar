package styles

import (
	"tipJar/globals/config"
	"tipJar/globals/log"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

var DefaultStyler *Styler

type Styler struct {
	tea.Model
	AccentColor   lg.Color
	InactiveColor lg.Color
	TextColor     lg.Color

	termWidth  int
	termHeight int
}

func NewStyler(config *config.Config) *Styler {
	log.Debug("creating new styler")
	return &Styler{
		AccentColor:   config.AccentColor,
		InactiveColor: config.InactiveColor,
		TextColor:     config.TextColor,
	}
}

func (s *Styler) Init() tea.Cmd {
	return nil
}

func (s *Styler) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Debug("updating window size", "width", msg.Width, "height", msg.Height)
		s.termWidth = msg.Width
		s.termHeight = msg.Height
	}
	return s, nil
}

func (s *Styler) View() string {
	return ""
}

func InitializeStyles(cfg *config.Config) {
	DefaultStyler = NewStyler(cfg)
}

func GetStyler() *Styler {
	if DefaultStyler == nil {
		panic("styles not initialized")
	}
	return DefaultStyler
}

// Styles

func (s *Styler) BaseStyle() lg.Style {
	return lg.NewStyle().
		Foreground(s.TextColor)
}

func (s *Styler) BorderStyle() lg.Style {
	return s.BaseStyle().
		Border(lg.RoundedBorder()).
		BorderForeground(s.InactiveColor)
}

func (s *Styler) DocStyle() lg.Style {
	return s.BorderStyle().
		Height(s.termHeight-2).
		Width(s.termWidth-2).
		MaxHeight(s.termHeight).
		MaxWidth(s.termWidth).
		Align(lg.Center, lg.Center)
}

func (s *Styler) PageStyle() lg.Style {
	return s.BorderStyle().
		Width(int(s.termWidth/5) * 4).
		Height(int(s.termHeight/3) * 2).
		Padding(4)
}
