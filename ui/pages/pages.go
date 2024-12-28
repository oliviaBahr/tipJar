package pages

import (
	"tipJar/core"
	"tipJar/ui/models"
	"tipJar/ui/pages/about"
	"tipJar/ui/pages/home"
	"tipJar/ui/pages/newtip"

	"github.com/charmbracelet/log"

	tea "github.com/charmbracelet/bubbletea"
)

type Pages []models.Page

func NewPages(jar *core.Jar) Pages {
	log.Debug("creating pages")
	return Pages{
		home.NewHomePage(jar),
		newtip.NewNewTipPage(),
		about.NewAboutPage(),
	}
}

func (p Pages) InitPages() []tea.Cmd {
	log.Debug("initializing pages")
	cmds := []tea.Cmd{}
	for i := range len(p) {
		log.Debug("initializing", "page", p[i].Title())
		cmds = append(cmds, p[i].Init())
	}
	log.Debug("initialized all pages")
	return cmds
}
