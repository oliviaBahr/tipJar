package home

import (
	"tipJar/core"

	"github.com/charmbracelet/log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type TipList struct {
	list list.Model
	jar  *core.Jar
}

func NewTipList(jar *core.Jar) *TipList {
	return &TipList{
		jar:  jar,
		list: list.New(nil, list.NewDefaultDelegate(), 0, 0),
	}
}

func (t TipList) Init() tea.Cmd {
	return nil
}

func (t TipList) View() string {
	log.Debug("rendering tip list")
	return t.list.View()
}

func (t TipList) Update(msg tea.Msg) (TipList, tea.Cmd) {
	var cmd tea.Cmd
	t.list, cmd = t.list.Update(msg)
	return t, cmd
}

func (t TipList) SetTips() TipList {
	items := []list.Item{}
	for _, tip := range t.jar.GetAllTips() {
		items = append(items, tip)
	}
	t.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
	return t
}
