package home

import (
	"tipJar/core"
	"tipJar/log"
	"tipJar/styles"
	"tipJar/ui/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HomePage struct {
	models.Page
	styler  *styles.Styler
	sidebar *Sidebar
	tipList *TipList
}

func NewHomePage(jar *core.Jar) HomePage {
	log.Debug("creating home page")
	return HomePage{
		styler:  jar.Styler,
		sidebar: NewSidebar(jar.Styler, []string{"tag1", "tag2"}),
		tipList: NewTipList(jar),
	}
}

func (p HomePage) Title() string {
	return "Home"
}

func (p HomePage) Init() tea.Cmd {
	return nil
}

func (p HomePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case TagToggled:
		// Here you can update the filtered tips based on the selected tags
		// For example:
		// p.filteredTips = p.jar.SearchByTags(p.sidebar.GetSelectedTags())
		return p, nil
	}

	var cmd tea.Cmd
	p.sidebar, cmd = p.sidebar.Update(msg)
	return p, cmd
}

func (p HomePage) View() string {
	log.Debug("rendering home page")
	return lipgloss.JoinHorizontal(lipgloss.Left, p.sidebar.View(), p.tipList.View())
}

func (p HomePage) PageStyle() lipgloss.Style {
	return p.styler.PageStyle()
}
