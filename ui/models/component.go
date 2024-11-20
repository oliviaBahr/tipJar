package models

import (
	"tipJar/globals/styles"

	tea "github.com/charmbracelet/bubbletea"
)

type BaseComponent interface {
	tea.Model
	Styler() *styles.Styler
}
