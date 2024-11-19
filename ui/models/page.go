package models

import (
	"github.com/charmbracelet/lipgloss"
)

type Page interface {
	BaseComponent
	Title() string
	PageStyle() lipgloss.Style
}
