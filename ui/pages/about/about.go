package about

import (
	"tipJar/ui/models"

	"github.com/charmbracelet/log"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type AboutPage struct {
	models.Page
	models.BaseComponent
}

func NewAboutPage() AboutPage {
	log.Debug("creating about page")
	return AboutPage{
		BaseComponent: models.NewBaseComponent(),
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

func (p AboutPage) PageStyle() lg.Style {
	return p.Styler.PageStyle()
}
