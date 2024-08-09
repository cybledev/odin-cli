package decorator

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	jsonData string
	done     bool
	spinner  spinner.Model
	viewport viewport.Model
	index    int
	dataChan chan string
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.fetchData(),
	)
}

func (m model) fetchData() tea.Cmd {
	return func() tea.Msg {
		data := <-m.dataChan
		return jsonMsg{data: data}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case jsonMsg:
		m.jsonData = msg.data
		m.done = true
		m.index = 0
		return m, animate()
	case errMsg:
		m.done = true
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case animateMsg:
		if m.index < len(m.jsonData) {
			m.index++
			return m, animate()
		}
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.done && m.index < len(m.jsonData) {
		return "\n" + colorizeJSON(m.jsonData[:m.index]) + "\n"
	} else if m.done {
		return "\n" + colorizeJSON(m.jsonData) + "\n"
	}
	return "\n" + m.spinner.View() + " Loading data...\n"
}

type animateMsg struct{}

func animate() tea.Cmd {
	return tea.Tick(time.Microsecond*100, func(time.Time) tea.Msg {
		return animateMsg{}
	})
}

func colorizeJSON(jsonData string) string {
	keyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	stringStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	numberStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	boolStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("5"))

	var result string
	inString := false
	inEscape := false

	for _, char := range jsonData {
		switch char {
		case '"':
			if !inEscape {
				inString = !inString
				result += stringStyle.Render(string(char))
			} else {
				result += stringStyle.Render(string(char))
			}
		case ':':
			if !inString {
				result += keyStyle.Render(string(char))
			} else {
				result += stringStyle.Render(string(char))
			}
		case '{', '}', '[', ']', ',':
			if !inString {
				result += string(char)
			} else {
				result += stringStyle.Render(string(char))
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if !inString {
				result += numberStyle.Render(string(char))
			} else {
				result += stringStyle.Render(string(char))
			}
		case 't', 'r', 'u', 'e', 'f', 'a', 'l', 's', 'n':
			if !inString {
				result += boolStyle.Render(string(char))
			} else {
				result += stringStyle.Render(string(char))
			}
		case '\\':
			inEscape = !inEscape
			result += stringStyle.Render(string(char))
		default:
			if !inString {
				result += string(char)
			} else {
				result += stringStyle.Render(string(char))
			}
		}

		if char != '\\' {
			inEscape = false
		}
	}

	return result
}

type jsonMsg struct {
	data string
}

type errMsg struct {
	err error
}
