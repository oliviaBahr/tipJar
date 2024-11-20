package ui

import (
	"os"
	"tipJar/core"
	"tipJar/globals/config"
	"tipJar/globals/log"
	"tipJar/ui/models"
	"tipJar/ui/pages"
	"tipJar/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type model struct {
	tea.Model
	models.BaseComponent
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
		BaseComponent: models.NewBaseComponent(),
		Jar:           jar,
		NavBar:        navBar,
		tabs:          tabs,
		err:           nil,
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
		default:
			m.NavBar.Update(msg)
		}
	case tea.WindowSizeMsg:
		m.Styler.Update(msg)
	}

	_, cmd := m.activeTab().Update(msg)
	return m, cmd
}

func (m model) View() string {
	log.Debug("rendering main view")
	pageV := m.activeTab().View()
	wholeV := lg.JoinVertical(lg.Center, m.NavBar.View(), pageV)
	return m.Styler.DocStyle().Render(wholeV)
}

func (m model) activeTab() models.Page {
	return m.tabs[m.NavBar.ActiveTab]
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
	if _, err := tea.NewProgram(m, tea.WithOutput(os.Stdout)).Run(); err != nil {
		log.Fatal("failed to run ui", "error", err)
	}
	return nil
}
