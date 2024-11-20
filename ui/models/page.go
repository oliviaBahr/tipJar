package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Page interface {
	tea.Model
	Title() string
	PageStyle() lipgloss.Style
}
