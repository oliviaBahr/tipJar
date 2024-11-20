package pages

import (
	"tipJar/core"
	"tipJar/globals/log"
	"tipJar/ui/models"
	"tipJar/ui/pages/about"
	"tipJar/ui/pages/home"
	"tipJar/ui/pages/newtip"

	tea "github.com/charmbracelet/bubbletea"
)

type Pages []models.Page

func NewPages(jar *core.Jar) Pages {
	log.Debug("creating pages")
	return Pages{
		home.NewHomePage(jar),
		newtip.NewNewTipPage(jar.Styler),
		about.NewAboutPage(jar.Styler),
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
