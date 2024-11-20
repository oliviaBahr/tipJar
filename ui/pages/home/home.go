package home

import (
	"tipJar/core"
	"tipJar/globals/log"
	"tipJar/ui/models"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type HomePage struct {
	models.Page
	models.BaseComponent
	sidebar *Sidebar
	tipList *TipList
}

func NewHomePage(jar *core.Jar) HomePage {
	log.Debug("creating home page")
	return HomePage{
		BaseComponent: models.NewBaseComponent(),
		sidebar:       NewSidebar([]string{"tag1", "tag2"}),
		tipList:       NewTipList(jar),
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
	v := lg.JoinHorizontal(lg.Left, p.sidebar.View(), p.tipList.View())
	return p.PageStyle().Render(v)
}

func (p HomePage) PageStyle() lg.Style {
	return p.Styler.PageStyle()
}
