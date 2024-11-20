package home

import (
	"tipJar/globals/log"
	"tipJar/globals/styles"

	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Sidebar struct {
	searchInput  string
	tags         []string
	selectedTags []string
	style        lipgloss.Style
	focused      bool
	styler       *styles.Styler
}

type TagToggled struct {
	Tag      string
	Selected bool
}

func NewSidebar(styler *styles.Styler, tags []string) *Sidebar {
	style := lipgloss.NewStyle().
		Inherit(styler.BorderStyle()).
		Width(16).
		Height(styler.PageStyle().GetHeight()-1).
		Border(lipgloss.RoundedBorder(), true, true, true, false).
		Align(lipgloss.Center)

	return &Sidebar{
		tags:         tags,
		selectedTags: []string{},
		style:        style,
		styler:       styler,
	}
}

func (s *Sidebar) Update(msg tea.Msg) (*Sidebar, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if s.focused {
				// Toggle tag selection when enter is pressed
				tag := s.tags[s.getSelectedTagIndex()]
				selected := s.toggleTag(tag)
				return s, func() tea.Msg {
					return TagToggled{Tag: tag, Selected: selected}
				}
			}
		case "up", "down":
			// Handle tag navigation
			if s.focused {
				s.handleTagNavigation(msg.String())
			}
		case "/":
			s.focused = true
		case "esc":
			s.focused = false
		default:
			if s.focused {
				s.handleSearchInput(msg.String())
			}
		}
	}
	return s, nil
}

func (s *Sidebar) View() string {
	log.Debug("rendering sidebar")
	// Build search bar
	searchStyle := lipgloss.NewStyle().
		Padding(0, 1).
		Inherit(s.styler.BorderStyle())

	searchPrompt := "🔍 "
	if s.focused {
		searchPrompt += s.searchInput + "█"
	} else {
		searchPrompt += "Search..."
	}

	searchBar := searchStyle.Render(searchPrompt)

	// Build tags list
	var tagList []string
	for _, tag := range s.tags {
		tagStyle := lipgloss.NewStyle()
		if s.isTagSelected(tag) {
			tagStyle = tagStyle.
				Background(s.styler.AccentColor).
				Foreground(lipgloss.NoColor{})
		}
		tagList = append(tagList, tagStyle.Render("# "+tag))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		searchBar,
		"",
		strings.Join(tagList, "\n"),
	)

	return s.style.Render(content)
}

func (s *Sidebar) toggleTag(tag string) bool {
	for i, selectedTag := range s.selectedTags {
		if selectedTag == tag {
			s.selectedTags = append(s.selectedTags[:i], s.selectedTags[i+1:]...)
			return false
		}
	}
	s.selectedTags = append(s.selectedTags, tag)
	return true
}

func (s *Sidebar) isTagSelected(tag string) bool {
	for _, selectedTag := range s.selectedTags {
		if selectedTag == tag {
			return true
		}
	}
	return false
}

func (s *Sidebar) handleSearchInput(key string) {
	if key == "backspace" && len(s.searchInput) > 0 {
		s.searchInput = s.searchInput[:len(s.searchInput)-1]
	} else if len(key) == 1 {
		s.searchInput += key
	}
}

func (s *Sidebar) handleTagNavigation(key string) {
	// Implement tag navigation logic
}

func (s *Sidebar) getSelectedTagIndex() int {
	// Implement selected tag index logic
	return 0
}

func (s *Sidebar) GetSelectedTags() []string {
	return s.selectedTags
}

func (s *Sidebar) GetSearchInput() string {
	return s.searchInput
}
