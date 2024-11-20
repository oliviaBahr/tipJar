package ui

import (
	"tipJar/core"
	"tipJar/globals/config"
	"tipJar/globals/log"
	"tipJar/ui/models"
	"tipJar/ui/pages"
	"tipJar/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	tea.Model
	Jar *core.Jar

	NavBar models.NavBar
	tabs   pages.Pages

	err error
}

func initialModel(jar *core.Jar) model {
	log.Debug("--- creating initial model ---")
	tabs := pages.NewPages(jar)
	navBar := models.NewNavBar(tabs)

	return model{
		Jar:    jar,
		NavBar: navBar,
		tabs:   tabs,
		err:    nil,
	}
}

func (m model) Init() tea.Cmd {
	log.Debug("model Init method")
	cmds := m.tabs.InitPages()
	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			log.Info("quitting")
			return m, tea.Quit
		}
	}

	m.NavBar, _ = m.NavBar.Update(msg)
	cmd := m.updateActiveTab(msg)
	return m, cmd
}

func (m model) View() string {
	log.Debug("rendering main view")
	pageV := m.activeTabView()
	wholeV := lipgloss.JoinVertical(lipgloss.Center, m.navBarView(), pageV)
	return wholeV
}

func (m model) activeTabView() string {
	page := m.tabs[m.NavBar.ActiveTab]
	log.Debug("rendering active tab", "page", page.Title())
	return page.View()
}

func (m model) updateActiveTab(msg tea.Msg) tea.Cmd {
	log.Debug("updating active tab", "page", m.tabs[m.NavBar.ActiveTab].Title())
	page, cmd := m.tabs[m.NavBar.ActiveTab].Update(msg)
	if p, ok := page.(models.Page); ok {
		m.tabs[m.NavBar.ActiveTab] = p
	}
	return cmd
}

func (m model) navBarView() string {
	log.Debug("rendering nav bar")
	return m.NavBar.View()
}

func RunUI() error {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config", "error", err)
	}

	// Initialize styles first
	styles.InitializeStyles(cfg)

	// Load jar
	jar, err := core.LoadJar(cfg)
	if err != nil {
		log.Fatal("failed to load jar", "error", err)
	}

	// Run ui
	m := initialModel(jar)

	log.Debug("running ui")
	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal("failed to run ui", "error", err)
	}
	return nil
}
