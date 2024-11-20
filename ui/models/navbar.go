package models

import (
	"tipJar/globals/log"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type NavBar struct {
	tea.Model
	BaseComponent
	Tabs      []string
	ActiveTab int

	numTabs int
}

func NewNavBar(pageList []Page) NavBar {
	tabs := []string{}
	for _, page := range pageList {
		tabs = append(tabs, page.Title())
	}

	return NavBar{
		BaseComponent: NewBaseComponent(),
		Tabs:          tabs,
		numTabs:       len(tabs),
	}
}

func (n NavBar) Init() tea.Cmd {
	return nil
}

func (n *NavBar) View() string {
	var renderedTabs []string
	for i, t := range n.Tabs {
		style := n.Styler.BorderStyle()
		if n.ActiveTab == i {
			style = style.BorderForeground(n.Styler.AccentColor)
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}
	joined := lg.JoinHorizontal(lg.Top, renderedTabs...)
	return joined
}

func (n *NavBar) Update(msg tea.Msg) (NavBar, tea.Cmd) {
	log.Debug("navbar update")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "right":
			log.Debug("navbar right")
			n.ActiveTab = (n.ActiveTab + 1) % n.numTabs
		case "left":
			log.Debug("navbar left")
			n.ActiveTab = (n.ActiveTab - 1 + n.numTabs) % n.numTabs
		}
	}
	return *n, nil
}

func (n *NavBar) Width() int {
	return lg.Width(n.View())
}
