package about

import (
	"tipJar/globals/log"
	"tipJar/globals/styles"
	"tipJar/ui/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AboutPage struct {
	models.Page
	styler *styles.Styler
}

func NewAboutPage(styler *styles.Styler) AboutPage {
	log.Debug("creating about page")
	return AboutPage{
		styler: styler,
	}
}

func (p AboutPage) Title() string {
	return "About"
}

func (p AboutPage) Init() tea.Cmd {
	log.Debug("initializing about page")
	return nil
}

func (p AboutPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

func (p AboutPage) View() string {
	return p.PageStyle().Render("About TipJar\n\n" +
		"A simple CLI tool to manage and track your tips.\n" +
		"Created with ❤️ using Go and Bubble Tea")
}

func (p AboutPage) PageStyle() lipgloss.Style {
	return p.styler.PageStyle()
}
