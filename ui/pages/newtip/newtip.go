package newtip

import (
	"tipJar/log"
	"tipJar/styles"
	"tipJar/ui/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type newTipPage struct {
	models.Page
	styler *styles.Styler
	form   *huh.Form
}

func NewNewTipPage(styler *styles.Styler) *newTipPage {
	log.Debug("creating new tip page")
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Title").
				Key("title"),
			huh.NewInput().
				Title("Description").
				Key("description"),
			huh.NewInput().
				Title("Links").
				Key("links"),
			huh.NewInput().
				Title("Tags").
				Key("tags"),
			huh.NewConfirm().
				Affirmative("Save").
				Negative("Cancel").
				Key("confirm"),
		),
	)

	return &newTipPage{
		styler: styler,
		form:   form,
	}
}

func (p newTipPage) Title() string {
	return "New Tip"
}

func (p newTipPage) Init() tea.Cmd {
	p.form.Init()
	return nil
}

func (p newTipPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var keyCmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			keyCmd = p.form.PrevField()
		case "down":
			keyCmd = p.form.NextField()
		}
	}
	form, cmd1 := p.form.Update(tea.Batch(keyCmd))
	form, cmd2 := form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		p.form = f
		cmds = append(cmds, cmd1, cmd2)
	}
	return p, tea.Batch(cmds...)
}

func (p newTipPage) View() string {
	return p.form.View()
}

func (p newTipPage) PageStyle() lipgloss.Style {
	return p.styler.PageStyle()
}
