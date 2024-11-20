package models

import (
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type Page interface {
	tea.Model
	Title() string
	PageStyle() lg.Style
}
