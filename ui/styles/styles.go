package styles

import (
	"os"
	"tipJar/globals/config"
	"tipJar/globals/log"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var DefaultStyler *Styler

type Styler struct {
	AccentColor   lipgloss.Color
	InactiveColor lipgloss.Color
	TextColor     lipgloss.Color
}

func NewStyler(config *config.Config) *Styler {
	log.Debug("creating new styler")
	aColor := config.AccentColor
	iColor := config.InactiveColor
	tColor := config.TextColor

	return &Styler{
		AccentColor:   aColor,
		InactiveColor: iColor,
		TextColor:     tColor,
	}
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

func (s *Styler) BorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(s.InactiveColor)
}

func (s *Styler) DocStyle() lipgloss.Style {
	return lipgloss.NewStyle().Padding(1).Align(lipgloss.Center)
}

func (s *Styler) PageStyle() lipgloss.Style {
	termWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	return s.BorderStyle().
		Padding(1).
		Width(termWidth - 20).
		Height(18).
		BorderForeground(s.AccentColor).
		Foreground(s.TextColor)
}
